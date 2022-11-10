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

var dsn = "root:helloboy@tcp(localhost:3309)/icrm?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB
var err error

func main() {
	gin.SetMode(gin.ReleaseMode)
	gob.Register(Employee{})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                              // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name

		}})
	//db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	//db.LogMode(true)
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
	router.GET("/products", allProducts)
	router.GET("/provinces", fetchProvines)
	router.GET("/city/:code", fetchChildCitys)
	router.GET("/topcity", ListTopAreas)
	router.GET("/childcity/:parentId", ListChildAreas)

	router.GET("/market/areas", ListAreas)
	router.GET("/market/provinces/:areaId", ListMarketProvinces)
	router.GET("/market/citys/:provinceId", ListMarketCitys)

	router.GET("/hospital/:hospitalId", fetchHospital) // handlerCustomer)
	router.POST("/hospitals/list", QueryHospitals)     // handlerCustomer)
	router.POST("/hospital/add", AddHospital)
	router.POST("/hospital/update", UpdateHospital)
	router.GET("/myHospitals", MyHospitals)

	router.GET("/customer/:customerId", fetchCustomer) // handlerCustomer)
	router.POST("/customers/list", QueryCustomers)     // handlerCustomer)
	router.POST("/customer/add", AddCustomer)
	router.POST("/customer/update", UpdateCustomer)
	router.GET("/myCustomers", MyCustomers)

	router.POST("/estimate/save", AddEstimate)

	router.Run(":3000")

}
