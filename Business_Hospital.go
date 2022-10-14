package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

type Hospital struct {
	ID    uint   `gorm:"column:ID"`
	Name  string `gorm:"column:name"`
	Code  string `gorm:"column:Code"`
	Grade string `gorm:"column:Grade"`
	HType string `gorm:"column:htype"`
}

func AddHospital(c *gin.Context) {
	var obj Hospital
	if err := c.BindJSON(&obj); err != nil {
		fmt.Println("发生错误")
		fmt.Println(obj)
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}

	if err := db.Create(&obj).Error; err != nil {
		fmt.Println("添加hospital到数据库失败：", err)
		return
	}

	fmt.Println(obj)
	c.JSON(http.StatusOK, obj)
}

/***************  单个客户信息  ***********************/
func fetchHospital(c *gin.Context) {

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
func UpdateHospital(c *gin.Context) {

	var obj Hospital
	if err := c.BindJSON(&obj); err != nil {
		fmt.Println("解析Hospital的json发生错误")
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

/***************  客户列表  ***********************/
func ListHospitals(c *gin.Context) {
	var results []map[string]interface{}
	var pagination Pagination
	var ct int64
	var hospital Hospital
	size, offset, count, sort := PaginationInf(c)
	pagination.PageSize = size
	if count == 0 {
		db.Model(&hospital).Count(&ct)
		pagination.TotalRows = ct
		pagination.TotalPages = int(math.Ceil(float64(ct) / float64(pagination.PageSize)))
	}

	err := db.Raw(`SELECT A.ID, A.Name,C6.name as City,C7.name as Province,
	C1.Label as HType,C4.label as Grade
	FROM Hospital A
	left join code C1 on A.htype  =  C1.code AND C1.codeType='HospitalType'
	left join code C4 on A.Grade  =  C4.code AND C4.codeType='HospitalGrade'
	left join city C6 ON A.Code   =  C6.code
	left join city C7 ON C7.code  =  C6.parentId order by ? limit ?,?
   `, sort, offset, size).Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	pagination.Rows = results
	c.JSON(http.StatusOK, pagination)

}

func QueryHospitals(c *gin.Context) {
	var objs []map[string]interface{}

	var pagination Pagination
	var ct int64
	var hospital Hospital
	size, offset, count, sort := PaginationInf(c)

	if count == 1 {
		db.Model(&hospital).Count(&ct)
		pagination.TotalRows = ct
		pagination.TotalPages = int(math.Ceil(float64(ct) / float64(pagination.PageSize)))
	}

	err := db.Raw(`SELECT A.ID, A.Name,C6.name as City,C7.name as Province,
	C1.Label as HType,C4.label as Grade
	FROM Hospital A
	left join code C1 on A.htype  =  C1.code AND C1.codeType='HospitalType'
	left join code C4 on A.Grade  =  C4.code AND C4.codeType='HospitalGrade'
	left join city C6 ON A.Code   =  C6.code
	left join city C7 ON C7.code  =  C6.parentId
	WHERE A.code IN ( WITH RECURSIVE CTE1 as
	(  select code from city where code IN (?)   UNION ALL
		select t1.code from city t1 inner join CTE1 t2  on t1.parentID = t2.code
	) SELECT * FROM CTE1)
  order by ? limit ?,? ;`, sort, offset, size).Find(&objs).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, objs)
}
