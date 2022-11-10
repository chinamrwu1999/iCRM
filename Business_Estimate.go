package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
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
    fmt.Println(objs)
	user := getUserInf(c)
	fmt.Println(user)
	for _, obj := range objs {

		(&obj).Submiter = user.ID

	}
	fmt.Println(objs)
	if err := db.Create(&objs).Error; err != nil {
		fmt.Println("添加hospital到数据库失败：", err)
		return
	}

	//fmt.Println(objs)
	c.JSON(http.StatusOK, objs)
}
