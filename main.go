package main

import (
	"os"

	"github.com/sourabhsd87/URL_Shortner/config"
	"github.com/sourabhsd87/URL_Shortner/db"
	"github.com/sourabhsd87/URL_Shortner/routes"
)

func main() {
	dir, _ := os.Getwd()
	config.LoadEnv()
	config.LoggerInit(dir)
	db.DbInit()
	config.InitRedis()
	config.InitOAuth()

	r := routes.SetupRouter()
	r.Run(config.Host + ":" + config.Port)
}

