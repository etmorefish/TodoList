package main

import (
	"todo-list/conf"
	"todo-list/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
