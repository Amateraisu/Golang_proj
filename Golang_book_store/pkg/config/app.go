package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

const (
	DBUsername = "KaiKaizxc"
	DBPassword = "DiaoChan"
	DBHost     = "localhost"
	DBPort     = "3390"
	DBName     = "LailaBase"
)

var DBURL = DBUsername + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	d, err := gorm.Open("mysql", "Kai: KaiKaizxc/KaiTable?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
