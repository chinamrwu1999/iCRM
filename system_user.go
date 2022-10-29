package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin" //  go get -u github.com/gin-gonic/gin
)

type Employee struct {
	ID           string // 主键
	Departmentid int8
	Name         string
	Email        string
	Password     string
	Gender       string
	Role         string
	Status       string
}

func UserLogin(c *gin.Context) {
	var params map[string]string

	if err := c.BindJSON(&params); err != nil {
		fmt.Println("发生错误")
		c.String(http.StatusBadRequest, "错误:%v", err)
		return
	}
	userId := params["UserId"]
	password := params["Password"]
	fmt.Printf("\nuserId=%s password=%s\n", userId, password)
	var user Employee

	err := db.Raw("SELECT * FROM employee WHERE ID=? and password=MD5(?)", userId, password).Find(&user).Error

	if err != nil {
		fmt.Println("user login error")
		c.JSON(http.StatusOK, Message{"error", "userId and password not match", ""})
	}

	session := sessions.Default(c)

	token := GetMD5Hash(userId + "@ms")
	session.Set(token, user)
	session.Save()
	c.Header("token", token)
	c.JSON(http.StatusOK, user)
}

func getUserInf(c *gin.Context) Employee {
	token := c.GetHeader("token")
	println(token)
	session := sessions.Default(c)
	user := session.Get(token).(Employee)
	return (user)

}
