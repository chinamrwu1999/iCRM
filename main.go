package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dsn = "root:helloboy@tcp(localhost:3306)/icrm?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB
var err error

/*
func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	// 分页功能实现
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(ctx.Query("page"))
		if page == 0 {
			page = 1
		}
		size, _ := strconv.Atoi(ctx.Query("size"))
		fmt.Println(size)
		switch{
		case size > 100:
			size = 100
		case size < 1:
			size = 50
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
*/

func PaginationInf(ctx *gin.Context) (int, int, int, string) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	count, _ := strconv.Atoi(ctx.Query("count"))
	sort := ctx.Query("sort")
	if page == 0 {
		page = 1
	}
	switch {
	case size > 100:
		size = 100
	case size < 1:
		size = 50
	}
	if sort == "" {
		sort = "ID"
	}
	offset := (page - 1) * size
	return size, offset, count, sort
}

func main() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                              // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name

		}})
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
		case "query":
			QueryHospitals(ctx)
		case "add":
			AddHospital(ctx)
		case "update":
			UpdateHospital(ctx)
		}
	}) // handlerCustomer)

	router.GET("/market/areas", ListAreas)
	router.GET("/market/provinces/:areaId", ListMarketProvinces)
	router.GET("/market/citys/:provinceId", ListMarketCitys)

	router.Run(":3000")

}
