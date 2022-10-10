package main

import "fmt"

func fetchCodes(codeType string) []Code {

	var objs []Code
	err := db.Where("codetype=?", codeType).Order("displayOrder").Find(&objs).Error
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(objs)
	return (objs)
}

func cityProvines() []City {
	var objs []City
	err := db.Raw("select code,name FROM citys where parentid is null order by code").Find(&objs).Error
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(objs)
	return (objs)
}

func cityChildren(code string) []City {
	var objs []City
	fmt.Println(code)
	err := db.Raw("select code,name FROM citys where parentid = ? order by code", code).Find(&objs).Error
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(objs)
	return (objs)
}

func fetchCustomerList() []Customer {
	var objs []Customer
	err := db.Raw(`SELECT A.ID, A.FullName,A.ShortName,A.Address,A.Description,C6.name as Province,C7.name as City,
	C1.Label as CType,C2.label as Status,C3.label as Scale,C4.label as level,C5.label as GetWay  
	FROM customers A
	left join codes C1 on A.ctype  =  C1.code
	left join codes C2 on A.status =  C2.code
	left join codes C3 on A.scale  =  C3.code
	left join codes C4 on A.level  =  C4.code
	left join codes C5 on A.getway =  C5.code
	left join citys C6 ON A.province= C6.code
	left join citys C7 ON A.city    = C7.code
    order by ID `).Find(&objs).Error
	if err != nil {
		fmt.Println(err)
		return (nil)
	}
	fmt.Println(objs)
	return (objs)
}
