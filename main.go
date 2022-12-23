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
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var dsn = "root:helloboy@tcp(localhost:3309)/icrm?charset=utf8mb4&parseTime=True&loc=Local"

// var dsn = "chinamrwu:helloboy@tcp(www.amswh.com:3306)/icrm?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB
var err error

func main() {
	gin.SetMode(gin.ReleaseMode)
	gob.Register(Employee{})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                              // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name

		}})
	//db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	//db.LogMode(true)
	router := gin.Default()
	// var htmlDir = "D:/tools/nginx-1.12.2/html/iCRM"
	var htmlDir = "C:/Tools/nginx-1.20.2/html/iCRM"
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob(htmlDir + "/*.html")
	router.StaticFS("/assets", http.Dir(htmlDir+"/assets"))
	router.StaticFile("/favicon.ico", htmlDir+"/favicon.ico")
	router.GET("/", index)
	router.POST("/login", UserLogin)
	router.POST("/goal/add", saveGoal)

	router.POST("/salesperson", SalesPersonByCity)

	router.GET("/codes/:type", getCodes)
	router.GET("/products", allProducts)
	router.GET("/provinces", fetchProvines)
	router.GET("/city/:code", fetchChildCitys)
	router.GET("/topcity", ListTopAreas)
	router.GET("/childcity/:parentId", ListChildAreas)

	router.GET("/market/areas", ListAreas)
	router.GET("/market/provinces/:areaId", ListMarketProvinces)
	router.GET("/market/citys/:provinceId", ListMarketCitys)

	router.GET("/hospital/:hospitalId", fetchHospital) //
	router.POST("/hospitals/list", QueryHospitals)     //
	router.POST("/hospital/add", AddHospital)
	router.POST("/hospital/update", UpdateHospital)
	router.GET("/myHospitals", MyHospitals)

	router.GET("/customer/:customerId", fetchCustomer) // handlerCustomer)
	router.POST("/customers/list", QueryCustomers)     //
	router.POST("/customer/add", AddCustomer)
	router.POST("/customer/update", UpdateCustomer)
	router.GET("/myCustomers", MyCustomers)

	router.GET("/proxy/:customerId", fetchProxy) //
	router.POST("/proxys/list", QueryProxys)     //
	router.POST("/proxy/add", AddProxy)
	router.POST("/proxy/update", UpdateCustomer)
	router.GET("/myProxys", MyProxys)

	router.POST("/EmployeeLogs", QueryEmployeeLogs) //
	router.POST("/QueryLogs", QueryLogs)            //
	router.POST("/Blogs/add", AddLogs)

	router.POST("/estimate/save", AddEstimate)
	router.POST("estimate/history", HistoryEstimate)

	router.Run(":3000")

}
