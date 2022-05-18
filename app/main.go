package main

import (
	"log"
	"netdisk-back/app/config"
	"netdisk-back/app/controller"
)

func main() {
	controller.Include()
	app := controller.Init()
	err := app.Run(":" + config.Get("server.port"))
	if err != nil {
		log.Println(err)
	}

}
