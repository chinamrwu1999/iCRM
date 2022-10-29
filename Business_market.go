package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

type MarketArea struct {
	AreaId uint   `gorm:"column:areaid"`
	Name   string `gorm:"column:name"`
}

type MarketProvince struct {
	AreaId uint   `gorm:"column:areaid"`
	Code   string `gorm:"column:code"`
}

type PositionName struct {
	Name string `gorm:"column:name"`
	Code string `gorm:"column:code"`
}

var SQL = `SELECT * FROM (SELECT * FROM city WHERE code IN ( 
	WITH RECURSIVE CTE1 as	(  
		 SELECT code FROM marketperson  WHERE employeeId=? UNION 
		 SELECT code FROM marketprovince WHERE areaId IN (SELECT code FROM marketperson WHERE employeeId=?)    
		 UNION ALL
			select t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
	) SELECT * FROM CTE1
	)) AS T0 `

func ListAreas(c *gin.Context) {
	var results []map[string]interface{}
	user := getUserInf(c)
	err := db.Raw(` SELECT areaID as Code,Name FROM marketarea WHERE areaId in (SELECT code FROM marketperson WHERE employeeId =?) UNION 
	SELECT Code,Name FROM city WHERE code IN (SELECT code FROM marketperson WHERE employeeId= ?) `, user.ID, user.ID).Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, results)

}

func ListTopAreas(c *gin.Context) {
	var results []map[string]interface{}
	user := getUserInf(c)
	err := db.Raw(` SELECT areaID as Code,Name FROM marketarea WHERE areaId in (SELECT code FROM marketperson WHERE employeeId =?) UNION 
	SELECT Code,Name FROM city WHERE code IN (SELECT code FROM marketperson WHERE employeeId= ?) `, user.ID, user.ID).Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, results)

}

func ListChildAreas(c *gin.Context) {
	var results []map[string]interface{}
	parentId := c.Param("parentId")
	err := db.Raw(` SELECT Code,Name FROM city WHERE code IN (SELECT code FROM marketprovince WHERE areaId=?) UNION
	SELECT Code,Name FROM city WHERE parentId=?  `, parentId, parentId).Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, results)

}

func ListMarketProvinces(c *gin.Context) {
	var results []map[string]interface{}
	areaId := c.Param("areaId")

	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, err)
	// }

	err = db.Raw(`SELECT A.*,B.Name FROM MarketProvince A LEFT JOIN city B ON A.Code=B.Code WHERE A.areaId=?`, areaId).Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, results)
}

func ListMarketCitys(c *gin.Context) {
	var results []map[string]interface{}
	provinceId, err := strconv.Atoi(c.Param("provinceId"))

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	err = db.Raw(`SELECT * FROM city A WHERE parentID =?`, provinceId).Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, results)
}
