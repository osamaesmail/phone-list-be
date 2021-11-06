package main

import (
	"fmt"
	"log"
	"phone-list/config"
	"phone-list/router"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Cfg().AppMode)
	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	app.Use(cors.New(corsConfig))

	router.InitRouter(app)

	if err := app.Run(fmt.Sprintf(":%d", config.Cfg().AppPort)); err != nil {
		panic(err)
	}

	log.Printf("[info] start http server listening localhost:%d", config.Cfg().AppPort)
}
