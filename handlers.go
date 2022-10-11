package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

/*
	func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		t, _ := template.ParseFiles("html/index.html")
		t.Execute(w, "HomePage")
	}

	func getPeoples(c *gin.Context) {
	    c.IndentedJSON(http.StatusOK, employees)
	}
*/
func saveGoal(c *gin.Context) {
	fmt.Println("receiving request for goal")
	//loc, _ := time.LoadLocation("Asia/Jakarta")
	var obj Goal
	//c.Bind(&obj)

	c.ShouldBindJSON(&obj)
	fmt.Println(obj.dueDate)
	fmt.Println(obj)
}



func getCodes(c *gin.Context) {
	codeType := c.Param("type")

	objs := fetchCodes(codeType)
	c.JSON(http.StatusOK, objs)
}

func fetchProvines(c *gin.Context) {
	objs := cityProvines()
	c.JSON(http.StatusOK, objs)
}

func fetchChildCitys(c *gin.Context) {
	code := c.Param("code")
	objs := cityChildren(code)
	c.JSON(http.StatusOK, objs)
}
