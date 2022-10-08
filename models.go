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

////////////////////////////////////////////////////

type Goal struct {
	ID         uint      `gorm:"column:ID"`
	Title      string    `gorm:"column:title"`
	dueDate    time.Time `gorm:"column:dueDate" time_format:"2019-01-02"`
	IsTeamGoal string
}

type Customer struct {
	ID          uint   `gorm:"column:ID"`
	FullName    string `gorm:"column:fullname"`
	ShortName   string `gorm:"column:shortname"`
	CType       string `gorm:"column:ctype"`
	Scale       string `gorm:"column:scale"`
	Status      string `gorm:"column:status"`
	Level       string `gorm:"column:level"`
	GetWay      string `gorm:"column:getway"`
	Nation      string `gorm:"column:nation"`
	Province    string `gorm:"column:province"`
	City        string `gorm:"column:city"`
	Address     string `gorm:"column:address"`
	Description string `gorm:"column:description"`
}

type Code struct {
	Label    string `gorm:"column:label"`
	Code     string `gorm:"column:value"`
	CodeType string `gorm:"column:codetype"`
	DisplayOrder string `gorm:"column:displayOrder"`
	Remark   string `gorm:"column:remark"`
}
