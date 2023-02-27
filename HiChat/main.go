package main

import (
	"HiChat/initialize"
	"HiChat/router"
)

func main() {
	//初始化数据库
	initialize.InitDB()
	//初始化日志
	initialize.InitLogger()

	router := router.Router()
	router.Run(":8000")
}
