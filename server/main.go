package main

import (
	"log"
	_ "server/docs"
	"server/routers/v1"
	"server/services/config"
	"server/services/database"
)

func init() {
	config.Setup()
	database.Setup()
}

// @title StreamSync API
// @version 1.0
// @BasePath /api/v1
func main() {
	conf := config.GetConfig()

	db := database.GetDB()
	server := v1.NewServer(db)

	err := server.Start(conf.Server.Address + ":" + conf.Server.Port)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

}
