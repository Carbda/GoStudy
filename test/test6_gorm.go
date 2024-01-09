package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	username := "root"
	password := "Wpp001029."
	host := "127.0.0.1"
	port := "3306"
	DBname := "go_study"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	//连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql.")
	} else {
		fmt.Println("connect success!")
	}
	funDB(db)
}

func funDB(db *gorm.DB) {

}
