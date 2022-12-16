package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

type Proxy struct {
	ID        uint   `gorm:"column:ID"`
	Name      string `gorm:"column:name"`
	ShortName string `gorm:"column:shortname"`
	Status    string `gorm:"column:status"`
	City      string `gorm:"column:city"`
	Address   string `gorm:"column:address"`
}

func AddProxy(c *gin.Context) {
	var obj Customer
	if err := c.BindJSON(&obj); err != nil {
		fmt.Println("发生错误")
		fmt.Println(obj)
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}

	if err := db.Create(&obj).Error; err != nil {
		fmt.Println("添加Goal到数据库失败：", err)
		return
	}

	fmt.Println(obj)
	c.JSON(http.StatusOK, obj)
}

/***************  单个客户信息  ***********************/
func fetchProxy(c *gin.Context) {

	proxyId, err := strconv.Atoi(c.Param("proxId"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	var obj Proxy
	err = db.First(&obj, proxyId).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, obj)
}

/***************  更新客户信息  ***********************/
func UpdateProxy(c *gin.Context) {

	var obj Proxy
	if err := c.BindJSON(&obj); err != nil {
		fmt.Println("解析Proxy的json发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	fmt.Println(obj)

	if err := db.Model(&obj).Updates(&obj).Error; err != nil {
		fmt.Println("更新Proxy到数据库失败：", err)
		return
	}

	fmt.Println(obj)
	c.JSON(http.StatusOK, obj)
}

/***************************** 查询  *****************************************/
func QueryProxys(c *gin.Context) {
	var objs []map[string]interface{}
	var paras map[string]string

	if err := c.BindJSON(&paras); err != nil {
		fmt.Println("发生错误:" + err.Error())
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	citys := paras["Citys"]
	name := paras["Txt"]

	var sql = `SELECT A.ID, A.Name,A.ShortName,A.Address,C6.name as City,C7.name as Province  
	FROM proxy A
	left join city C6 ON A.City   =  C6.code
	left join city C7 ON C7.code  =  C6.parentId 
	WHERE A.City IN (
		WITH RECURSIVE CTE1 as	(  
						SELECT code FROM marketperson  WHERE employeeId=? UNION 
						SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=? )    
						UNION ALL SELECT t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
		) SELECT * FROM CTE1
     )	`

	var pagination Pagination
	var ct int64
	size, offset, count, sort := PaginationInf(c)
	//fmt.Printf("offset=%d\n", offset)
	pagination.PageSize = size
	pagination.StartIndex = offset

	user := getUserInf(c)

	var err error
	if name != "" { // 模糊查询
		sql += "AND A.FullName like ? ORDER BY ? limit ?,?"
		err = db.Raw(sql, user.ID, user.ID, "%"+name+"%", sort, offset, size).Find(&objs).Error
	} else if citys != "" { //根据区域查询
		sql += `  AND A.City IN ( 
			WITH RECURSIVE CTE1 as	(  
				 SELECT code FROM marketprovince WHERE areaId =? UNION
				 SELECT ?  
				 UNION ALL
					select t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
			) SELECT * FROM CTE1
			)  ORDER BY ? limit ?,?`
		err = db.Raw(sql, user.ID, user.ID, citys, citys, sort, offset, size).Find(&objs).Error

	} else { //列出所有
		sql += " ORDER BY ? limit ?,?"
		err = db.Raw(sql, user.ID, user.ID, sort, offset, size).Find(&objs).Error
	}
	pagination.Rows = objs

	if count == 0 {
		countSQL := ` SELECT count(*) FROM proxy A
		WHERE A.city IN (
			WITH RECURSIVE CTE1 as	(  
							SELECT code FROM marketperson  WHERE employeeId=? UNION 
							SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=? )    
							UNION ALL SELECT t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
			) SELECT * FROM CTE1 ) `
		if name != "" {
			db.Raw(countSQL+` AND  A.FullName like ?`, user.ID, user.ID, "%"+name+"%").Count(&ct)
		} else if citys != "" {
			//var arr = strings.Split(citys, ",")
			db.Raw(countSQL+` AND A.city IN ( 
				WITH RECURSIVE CTE2 as	(  
					 SELECT code FROM marketprovince WHERE areaId =? UNION
					 SELECT ?  
					 UNION ALL
						select t1.code from city t1 inner join CTE2 t2  on t1.parentID = t2.code
				) SELECT * FROM CTE2) `, user.ID, user.ID, citys, citys).Count(&ct)
		} else {
			db.Raw(countSQL).Count(&ct)
		}
		pagination.StartIndex = 0
		pagination.TotalRows = ct
		pagination.TotalPages = int(math.Ceil(float64(ct) / float64(pagination.PageSize)))
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, pagination)
}

// 用户负责区域的客户列表
func MyProxys(c *gin.Context) {

	var pagination Pagination
	var ct int64
	size, offset, count, sort := PaginationInf(c)
	pagination.PageSize = size
	pagination.StartIndex = offset
	//var err error

	user := getUserInf(c)

	var sql = `SELECT A.ID, A.Name,A.ShortName,A.Address,C6.name as Province,C7.name as City
	FROM Proxy A
	left join city C6 ON A.City   =  C6.code
	left join city C7 ON C7.code  =  C6.parentId 
	WHERE A.City IN (
		WITH RECURSIVE CTE1 as	(  
						SELECT code FROM marketperson  WHERE employeeId=? UNION 
						SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=? )    
						UNION ALL SELECT t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
		) SELECT * FROM CTE1
     )	ORDER BY ? limit ?,?`
	var objs []map[string]interface{}
	err = db.Raw(sql, user.ID, user.ID, sort, offset, size).Find(&objs).Error
	pagination.Rows = objs

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}

	if count == 0 {
		err = db.Raw(` SELECT count(*) FROM proxy A where A.city in (WITH RECURSIVE CTE1 as	(  
						SELECT code FROM marketperson  WHERE employeeId=? UNION 
						SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=? )    
						UNION ALL
						   select t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
				   ) SELECT * FROM CTE1 )`, user.ID, user.ID).Count(&ct).Error

	}
	pagination.StartIndex = 0
	pagination.TotalRows = ct
	pagination.TotalPages = int(math.Ceil(float64(ct) / float64(pagination.PageSize)))

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}

	c.JSON(http.StatusOK, pagination)

}
