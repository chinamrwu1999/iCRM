package main

import "fmt"

func fetchCodes(codeType string) []Code {

	var objs []Code
	err := db.Where("codetype=?", codeType).Order("displayOrder").Find(&objs).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(objs)
	return (objs)
}
