package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
	"gorm.io/gorm/clause"
)

type Estimate struct {
	ProductId  string  `gorm:"column:productId"`
	CustomerId uint    `gorm:"column:customerId"`
	Year       uint    `gorm:"column:saleYear"`
	Month      uint    `gorm:"column:saleMonth"`
	Amount     float32 `gorm:"column:amount"`
	Price      float32 `gorm:"column:price"`
	Submiter   string  `gorm:"column:submiter"`
}

func AddEstimate(c *gin.Context) {
	objs := []Estimate{}
	if err := c.ShouldBindJSON(&objs); err != nil {
		fmt.Println("保存销售预估时，解析参数发生错误")
		fmt.Println(objs)
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	//fmt.Println(objs)
	user := getUserInf(c)
	//fmt.Println(user)
	for _, obj := range objs {

		(&obj).Submiter = user.ID

	}
	//fmt.Println(objs)

	if err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&objs).Error; err != nil {
		fmt.Println("添加hospital到数据库失败：", err)
		return
	}

	//fmt.Println(objs)
	c.JSON(http.StatusOK, objs)
}

func HistoryEstimate(c *gin.Context) {
	var params map[string]string

	if err := c.BindJSON(&params); err != nil {
		fmt.Println("发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	customerId, _ := strconv.Atoi(params["CustomerId"])
	productId := params["ProductId"]
	year, _ := strconv.Atoi(params["Year"])

	fmt.Printf("\ncustomerIdId=%d password=%s year=%d\n", customerId, productId, year)

	objs := []Estimate{}
	//var err error
	err = db.Raw("SELECT * FROM estimate WHERE customerId=? and productId=? and saleYear=?", customerId, productId, year).Find(&objs).Error
	if err != nil {
		fmt.Println("查询预估时候，发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	c.JSON(http.StatusOK, objs)
}
