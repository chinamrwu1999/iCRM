package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

func CustomerAdd(c *gin.Context) {
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

func CustomerList(c *gin.Context) {
	var objs []Customer = fetchCustomerList()

	c.JSON(http.StatusOK, objs)
}
