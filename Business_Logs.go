package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

type BusinessLog struct {
	ID         uint   `gorm:"column:ID"`
	EmployeeId string `gorm:"column:employeeId"`
	HospitalId int    `gorm:"column:hospitalId"`
	ProxyId    int    `gorm:"column:proxyId"`
	CustomerId int    `gorm:"column:customerId"`
	Content    string `gorm:"column:content"`
}

func AddLogs(c *gin.Context) {
	var obj BusinessLog
	if err := c.BindJSON(&obj); err != nil {
		fmt.Println("添工作日志时，解析参数发生错误")
		fmt.Println(obj)
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}

	if err := db.Create(&obj).Error; err != nil {
		fmt.Println("添加工作日志到数据库失败：", err)
		return
	}

	fmt.Println(obj)
	c.JSON(http.StatusOK, obj)
}

/***************  单个客户信息  ***********************/
func fetchLog(c *gin.Context) {

	hospitalId, err := strconv.Atoi(c.Param("hospitalId"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	var obj Hospital
	err = db.First(&obj, hospitalId).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, obj)
}

/***************  更新客户信息  ***********************/
func UpdateLog(c *gin.Context) {

	var obj Hospital
	if err := c.BindJSON(&obj); err != nil {
		fmt.Println("更新医院信息：解析Hospital的json发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	fmt.Println(obj)

	if err := db.Model(&obj).Updates(&obj).Error; err != nil {
		fmt.Println("更新Hospital到数据库失败：", err)
		return
	}

	fmt.Println(obj)
	c.JSON(http.StatusOK, obj)
}

/***************************** 查询  *****************************************/
func QueryLogs(c *gin.Context) {
	var objs []map[string]interface{}
	var paras map[string]string

	if err := c.BindJSON(&paras); err != nil {
		fmt.Println("查询医院：解析参数发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	citys := paras["Citys"]
	name := paras["Txt"]

	var sql = `SELECT A.ID, A.Name,C6.name as City,C7.name as Province,
	C1.Label as HType,C4.label as Grade
	FROM Hospital A
	left join code C1 ON A.htype  =  C1.code AND C1.codeType='HospitalType'
	left join code C4 ON A.Grade  =  C4.code AND C4.codeType='HospitalGrade'
	left join city C6 ON A.city   =  C6.code
	left join city C7 ON C7.code  =  C6.parentId 
	WHERE A.city IN (
		WITH RECURSIVE CTE1 as	(  
						SELECT code FROM marketperson  WHERE employeeId=? UNION 
						SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=? )    
						UNION ALL SELECT t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
		) SELECT * FROM CTE1
     )
	`
	user := getUserInf(c)
	var pagination Pagination
	var ct int64
	size, offset, count, sort := PaginationInf(c)
	//fmt.Printf("offset=%d\n", offset)
	pagination.PageSize = size
	pagination.StartIndex = offset
	var err error
	if name != "" { // 模糊查询
		sql += "AND A.name like ? ORDER BY ? limit ?,?"
		err = db.Raw(sql, user.ID, user.ID, "%"+name+"%", sort, offset, size).Find(&objs).Error
	} else if citys != "" { //根据区域查询
		//var arr = strings.Split(citys, ",")
		sql += `  AND A.city IN ( 
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
		countSQL := ` SELECT count(*) FROM hospital A
		WHERE A.city IN (
			WITH RECURSIVE CTE1 as	(  
							SELECT code FROM marketperson  WHERE employeeId=? UNION 
							SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=? )    
							UNION ALL SELECT t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
			) SELECT * FROM CTE1 ) `
		if name != "" {
			db.Raw(countSQL+` AND  A.name like ?`, user.ID, user.ID, "%"+name+"%").Count(&ct)
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

// 用户负责区域的医院列表
func MyBusinessLogs(c *gin.Context) {

	var pagination Pagination
	var ct int64
	size, offset, count, sort := PaginationInf(c)
	pagination.PageSize = size
	pagination.StartIndex = offset
	//var err error

	user := getUserInf(c)

	sql := `SELECT A.ID, A.Name,C6.name as City,C7.name as Province, 
    C1.Label as HType,C4.label as Grade 
    FROM Hospital A
    left join code C1 on A.htype  =  C1.code AND C1.codeType='HospitalType'
    left join code C4 on A.Grade  =  C4.code AND C4.codeType='HospitalGrade'
    left join city C6 ON A.City   =  C6.code
    left join city C7 ON C7.code  =  C6.parentId 
	WHERE A.city IN (
 	WITH RECURSIVE CTE1 AS (  
		SELECT distinct code from city where code IN (
			SELECT T2.code FROM marketperson T1,MarketProvince T2 where T1.code =T2.areaID AND T1.employeeID=?
		    UNION
			SELECT code FROM marketperson  where employeeID=?
  )UNION ALL   SELECT t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
	   ) SELECT * FROM CTE1    ) ORDER BY ? limit ?,?`

	var objs []map[string]interface{}
	err = db.Raw(sql, user.ID, user.ID, sort, offset, size).Find(&objs).Error
	pagination.Rows = objs

	if count == 0 {
		db.Raw(` SELECT count(*) FROM hospital A where A.city in (WITH RECURSIVE CTE1 as	(  
						SELECT code FROM marketperson  WHERE employeeId=? UNION 
						SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=? )    
						UNION ALL
						   select t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
				   ) SELECT * FROM CTE1 )`, user.ID, user.ID).Count(&ct)

	}
	pagination.StartIndex = 0
	pagination.TotalRows = ct
	pagination.TotalPages = int(math.Ceil(float64(ct) / float64(pagination.PageSize)))

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, pagination)

}
