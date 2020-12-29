package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (base *Controller) Dashboard(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	c.JSON(http.StatusOK, gin.H{"Welcome: ": user})
}
