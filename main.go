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

	router.GET("/codes/:type", getCodes)
	router.GET("/provinces", fetchProvines)
	router.GET("/city/:code", fetchChildCitys)
	//router.POST("/customer/add", CustomerAdd)
	//router.GET("/customer/list", handlerCustomerList)
	router.GET("/customer/:customerId", func(ctx *gin.Context) {
		switch ctx.Param("customerId") {
		case "list":
			ListCustomers(ctx)
		default:
			fetchCustomer(ctx)
		}
	}) // handlerCustomer)
	router.POST("/customer/:customerId", func(ctx *gin.Context) {
		switch ctx.Param("customerId") {
		case "add":
			AddCustomer(ctx)
		case "update":
			UpdateCustomer(ctx)
		}
	}) // handlerCustomer)

	router.GET("/hospital/:hospitalId", func(ctx *gin.Context) {
		switch ctx.Param("hospitalId") {
		case "list":
			ListHospitals(ctx)
		default:
			fetchHospital(ctx)
		}
	}) // handlerCustomer)
	router.POST("/hospital/:hospitalId", func(ctx *gin.Context) {
		switch ctx.Param("hospitalId") {
		case "add":
			AddHospital(ctx)
		case "update":
			UpdateHospital(ctx)
		}
	}) // handlerCustomer)
	router.Run(":3000")

}
