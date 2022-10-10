package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:helloboy@tcp(localhost:3306)/icrm?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")
	router.StaticFS("/assets", http.Dir("html/assets"))
	router.StaticFile("/favicon.ico", "html/favicon.ico")
	router.GET("/", index)
	router.POST("/goal/add", saveGoal)
	router.POST("/customer/add", CustomerAdd)
	router.GET("/customer/list", CustomerList)

	router.GET("/codes/:type", getCodes)
	router.GET("/provinces", fetchProvines)
	router.GET("/city/:code", fetchChildCitys)
	router.Run(":3000")

}
