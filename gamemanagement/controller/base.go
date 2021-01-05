package controller

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"fmt"
)

var (
	db *gorm.DB
	dbErr error
)

func init() {
	dsn := "g4m3m4n4g3m3nt:k1!vapi!m4n4g3M3Nt!@tcp(172.18.0.3:3306)/game"
	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		fmt.Println("Error.")
	}
	db.Debug()
}

func GetDB() *gorm.DB {
	return db
}
