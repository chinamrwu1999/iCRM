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
	userId := params["userId"]
	password := params["password"]

	var user Employee

	err := db.Raw("SELECT * FROM employee WHERE userID=? and password=MD5(?)", userId, password).Find(&user).Error

	if err != nil {
		c.JSON(http.StatusOK, Message{"error", "userId and password not match", ""})
	}

	session := sessions.Default(c)
	session.Set("userId", user)
	session.Save()
	fmt.Println(userId + "login successfully")
	c.JSON(http.StatusOK, Message{"success", "OK", ""})
}
