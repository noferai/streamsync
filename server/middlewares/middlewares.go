package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/config"
)

//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Conf.GetString("access_control.allow_origin"))
		c.Writer.Header().Set("Access-Control-Max-Age", config.Conf.GetString("access_control.max_age"))
		c.Writer.Header().Set("Access-Control-Allow-Methods", config.Conf.GetString("access_control.allow_methods"))
		c.Writer.Header().Set("Access-Control-Allow-Headers", config.Conf.GetString("access_control.allow_headers"))
		c.Writer.Header().Set("Access-Control-Expose-Headers", config.Conf.GetString("access_control.expose_headers"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", config.Conf.GetString("access_control.allow_credentials"))

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

//Generate a unique ID and attach it to each request for future reference or use
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", _uuid.String())
		c.Next()
	}
}
