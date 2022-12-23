package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

type Employee struct {
	ID       string `gorm:"column:ID"`
	Name     string `gorm:"column:name"`
	Role     string `gorm:"column:role"`
	Password string `gorm:"column:password"`
}

func UserLogin(c *gin.Context) {
	var params map[string]string

	if err := c.BindJSON(&params); err != nil {
		fmt.Println("发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	userId := params["UserId"]
	password := params["Password"]
	fmt.Printf("\nuserId=%s password=%s\n", userId, password)
	var user Employee

	err := db.Raw("SELECT * FROM employee WHERE ID=? and password=MD5(?)", userId, password).Find(&user).Error

	if err != nil {
		fmt.Println("user login error")
		c.JSON(http.StatusOK, Message{"error", "userId and password not match", ""})
	}

	session := sessions.Default(c)

	token := GetMD5Hash(userId + "@ms")
	session.Set(token, user)
	session.Save()
	c.Header("token", token)
	c.JSON(http.StatusOK, user)
}

func getUserInf(c *gin.Context) Employee {
	token := c.GetHeader("token")
	session := sessions.Default(c)
	user := session.Get(token).(Employee)
	return (user)

}

/***************************** 查询  *****************************************/
func SalesPersonByCity(c *gin.Context) {
	var objs []map[string]interface{}
	var paras map[string]string

	if err := c.BindJSON(&paras); err != nil {
		fmt.Println("查询销售：解析参数发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	citys := paras["Citys"]
	var sql = `SELECT A.ID, A.Name FROM Employee A
	WHERE A.ID IN ( 
	SELECT distinct employeeID FROM marketperson 
	WHERE code IN 
	(WITH RECURSIVE CTE1 AS	(  
		SELECT code FROM marketperson  WHERE employeeId=? UNION 
		SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=? )    
		UNION ALL SELECT t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
			) SELECT * FROM CTE1)
	) 
	`
	user := getUserInf(c)
	var pagination Pagination
	size, offset, _, _ := PaginationInf(c)
	//fmt.Printf("offset=%d\n", offset)
	pagination.PageSize = size
	pagination.StartIndex = offset
	var err error
	if citys != "" { //根据区域查询
		//var arr = strings.Split(citys, ",")
		sql += ` AND A.ID IN (
			SELECT distinct employeeID FROM marketperson 
			WHERE code IN (WITH RECURSIVE CTE1 AS (  
				 SELECT code FROM marketprovince WHERE areaId =? UNION
				 SELECT ?  
				 UNION ALL SELECT t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
			   ) SELECT * FROM CTE1 )
			) ORDER BY A.ID 
		`
		err = db.Raw(sql, user.ID, user.ID, citys, citys).Find(&objs).Error
	} else { //列出所有
		sql += " ORDER BY A.ID "
		err = db.Raw(sql, user.ID, user.ID).Find(&objs).Error
	}
	pagination.Rows = objs
	pagination.StartIndex = 0
	pagination.TotalRows = 0
	pagination.TotalPages = 1

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, pagination)
}
