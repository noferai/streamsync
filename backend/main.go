package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"chat/pkg/websocket"
)

// get heroku dynamic port from env
func getPort() string {
	if value, ok := os.LookupEnv("PORT"); ok {
		return value
	}
	return "8080"
}

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func main() {
	fmt.Println("App v0.03")
	r := gin.Default()
	pool := websocket.NewPool()

	go pool.Start()

	r.Use(static.Serve("/", static.LocalFile("./static", true)))
	r.GET("/ws", func(c *gin.Context) {
		serveWs(pool, c.Writer, c.Request)
	})

	r.Run(":" + getPort())

}
