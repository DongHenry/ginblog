package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	// 数据库
	model.InitDb()

	routes.InitRouter()
}
