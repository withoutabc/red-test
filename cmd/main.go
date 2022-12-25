package main

import (
	"proj/api"
	"proj/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
