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

func ListAreas(c *gin.Context) {
	var results []MarketArea

	err := db.Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, results)

}

func ListMarketProvinces(c *gin.Context) {
	var results []map[string]interface{}
	areaId, err := strconv.Atoi(c.Param("areaId"))

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

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
