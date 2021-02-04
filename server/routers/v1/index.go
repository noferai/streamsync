package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/middlewares"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{db: db}
	engine := gin.Default()

	// Middlewares
	engine.Use(middlewares.CORS())

	// Non-protected routes
	v1 := engine.Group("/api/v1")
	{
		rooms := v1.Group("/rooms")
		{
			rooms.GET("/", server.GetRooms)
			rooms.GET("/:id", server.GetRoom)
			rooms.POST("/create", server.CreateRoom)
			rooms.PUT("/update/:id", server.UpdateRoom)
			rooms.DELETE("delete/:id", server.DeleteRoom)
		}
	}

	server.router = engine
	return server
}

func (server *Server) Start(address string) error {
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return server.router.Run(address)
}
