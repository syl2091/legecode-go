package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/library?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}
