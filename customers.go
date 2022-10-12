package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)


type Customer struct {
	ID          uint   `gorm:"column:ID"`
	FullName    string `gorm:"column:fullname"`
	ShortName   string `gorm:"column:shortname"`
	CType       string `gorm:"column:ctype"`
	Scale       string `gorm:"column:scale"`
	Status      string `gorm:"column:status"`
	Level       string `gorm:"column:level"`
	GetWay      string `gorm:"column:getway"`
	Nation      string `gorm:"column:nation;default:'cn'"`
	Province    string `gorm:"column:province"`
	City        string `gorm:"column:city"`
	Address     string `gorm:"column:address"`
	Description string `gorm:"column:description"`
}

func AddCustomer(c *gin.Context) {
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
func fetchCustomer(c *gin.Context) {

	customerId, err := strconv.Atoi(c.Param("customerId"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	var obj Customer
	err = db.First(&obj, customerId).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, obj)
}

/***************  更新客户信息  ***********************/
func UpdateCustomer(c *gin.Context) {

	var obj Customer
	if err := c.BindJSON(&obj); err != nil {
		fmt.Println("解析Customer的json发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	fmt.Println(obj)

	if err := db.Model(&obj).Updates(&obj).Error; err != nil {
		fmt.Println("更新Customer到数据库失败：", err)
		return
	}

	fmt.Println(obj)
	c.JSON(http.StatusOK, obj)
}

/***************  客户列表  ***********************/
func ListCustomers(c *gin.Context) {
	var objs []Customer
	err := db.Raw(`SELECT A.ID, A.FullName,A.ShortName,A.Address,A.Description,C6.name as Province,C7.name as City,
	C1.Label as CType,C2.label as Status,C3.label as Scale,C4.label as level,C5.label as GetWay  
	FROM customers A
	left join codes C1 on A.ctype  =  C1.code AND C1.codeType='customerType'
	left join codes C2 on A.status =  C2.code AND C2.codeType='customerStatus'
	left join codes C3 on A.scale  =  C3.code AND C3.codeType='scale'
	left join codes C4 on A.level  =  C4.code AND C4.codeType='customerGrade'
	left join codes C5 on A.getway =  C5.code AND C5.codeType='customerWay'
	left join citys C6 ON A.province= C6.code 
	left join citys C7 ON A.city    = C7.code
    order by ID `).Find(&objs).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, objs)
}

func QueryCustomers(c *gin.Context) {
	var objs []Customer
	err := db.Raw(`SELECT A.ID, A.FullName,A.ShortName,A.Address,A.Description,C6.name as Province,C7.name as City,
	C1.Label as CType,C2.label as Status,C3.label as Scale,C4.label as level,C5.label as GetWay  
	FROM customers A
	left join codes C1 on A.ctype  =  C1.code AND C1.codeType='customerType'
	left join codes C2 on A.status =  C2.code AND C2.codeType='customerStatus'
	left join codes C3 on A.scale  =  C3.code AND C3.codeType='scale'
	left join codes C4 on A.level  =  C4.code AND C4.codeType='customerGrade'
	left join codes C5 on A.getway =  C5.code AND C5.codeType='customerWay'
	left join citys C6 ON A.province= C6.code 
	left join citys C7 ON A.city    = C7.code
    order by ID `).Find(&objs).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, objs)
}
