package api

import (
	"github.com/jinzhu/gorm"
	"server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	c := Controller{DB: db}

	// Non-protected routes
	v1 := r.Group("/api/v1")
	{
		rooms := v1.Group("/rooms")
		{
			rooms.GET("/", c.GetRooms)
			rooms.GET("/:id", c.GetRoom)
			rooms.POST("/create", c.CreateRoom)
			rooms.PUT("/update/:id", c.UpdateRoom)
			rooms.DELETE("delete/:id", c.DeleteRoom)
		}
	}

	// Protected routes
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"username1": "password1",
		"username2": "password2",
	}))

	// admin/dashboard endpoint
	authorized.GET("/dashboard", c.Dashboard)

	return r
}
