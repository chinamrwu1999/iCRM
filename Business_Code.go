package main

import "fmt"

func fetchCodes(codeType string) []Code {

	var objs []Code
	err := db.Where("codetype=?", codeType).Order("displayOrder").Find(&objs).Error
	if err != nil {
		fmt.Println(err)
	}
	return (objs)
}

func cityProvines() []City {
	var objs []City
	err := db.Raw("select code,name FROM city where parentid is null order by code").Find(&objs).Error
	if err != nil {
		fmt.Println(err)
	}
	return (objs)
}

func cityChildren(code string) []City {
	var objs []City
	err := db.Raw("select code,name FROM city where parentid = ? order by code", code).Find(&objs).Error
	if err != nil {
		fmt.Println(err)
	}
	return (objs)
}

func fetchProducts() []Product {

	var products []Product
	err := db.Find(&products).Error
	if err != nil {
		fmt.Println(err)
	}
	return (products)
}
