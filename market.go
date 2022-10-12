package main

type MarketArea struct{
    AreaId uint  `gorm:"column:areaid"`
	Name string   `gorm:"column:name"`
}