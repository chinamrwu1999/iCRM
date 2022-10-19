package main

import (
	"time"
)

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql

type Department struct {
	ID       uint   `gorm:"column:ID"`
	Name     string `gorm:"column:name"`
	ParentId uint   `gorm:"column:parentId"`
}



////////////////////////////////////////////////////

type Goal struct {
	ID         uint      `gorm:"column:ID"`
	Title      string    `gorm:"column:title"`
	dueDate    time.Time `gorm:"column:dueDate" time_format:"2019-01-02"`
	IsTeamGoal string
}

type Code struct {
	Label        string `gorm:"column:label"`
	Code         string `gorm:"column:value"`
	CodeType     string `gorm:"column:codetype"`
	DisplayOrder string `gorm:"column:displayOrder"`
	Remark       string `gorm:"column:remark"`
}

type City struct {
	Code string `gorm:"column:code"`
	Name string `gorm:"column:name"`
	//ParentId string `gorm:"column:parentId"`
}

type PostParameters struct {
	Codes string
	Txt   string
}

type Message struct {
	Type string
	content interface{}
	action string 
}


