package main

import (
	"santri-koding-trial/backend-golang/config"
	"santri-koding-trial/backend-golang/database"
	"santri-koding-trial/backend-golang/routes"
)

func main() {
	config.LoadEnv()
	database.InitDB()
	r := routes.SetupRouter()
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
