package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "server/docs"

	"server/api"
	"server/pkg/config"
	"server/pkg/database"
)

func init() {
	config.Setup()
	database.Setup()
}

// @title StreamSync API
// @version 1.0
// @BasePath /api/v1
func main() {
	configuration := config.GetConfig()

	db := database.GetDB()
	r := api.Setup(db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("127.0.0.1:" + configuration.Server.Port)
}
