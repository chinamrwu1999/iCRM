package main

import (
	"net/http"
	"strings"

	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dsn = "root:helloboy@tcp(localhost:3306)/icrm?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB
var err error

func main() {
	gob.Register(Employee{})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                              // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name

		}})
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("html/*.html")
	router.StaticFS("/assets", http.Dir("html/assets"))
	router.StaticFile("/favicon.ico", "html/favicon.ico")
	router.GET("/", index)
	router.POST("/login", UserLogin)
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

	router.GET("/hospital/:hospitalId", fetchHospital) // handlerCustomer)
	router.POST("/hospitals/list", QueryHospitals)     // handlerCustomer)
	router.POST("/hospital/add", AddHospital)
	router.POST("/hospital/update", UpdateHospital)
	router.GET("/market/areas", ListAreas)
	router.GET("/market/provinces/:areaId", ListMarketProvinces)
	router.GET("/market/citys/:provinceId", ListMarketCitys)

	router.Run(":3000")

}
