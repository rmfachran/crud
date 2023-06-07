package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GormMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/bootcamp_crud?charset=utf8mb4&parseTime=True&loc=UTC"), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}
