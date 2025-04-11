package main

import (
	"github.com/yuhari7/backend_supervision/api"
	"github.com/yuhari7/backend_supervision/config"
)

func main() {
	config.InitDB()

	e := api.NewServer()
	e.Logger.Fatal(e.Start(":8080"))
}
