package main

import (
	"net/http"
	"os"
	"fmt"
	"com.werner/gotchat/resources"
	"github.com/apsdehal/go-logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var connections map[*websocket.Conn]bool;

func main() {
	connections = make(map[*websocket.Conn]bool)
	router := gin.Default()
	router.LoadHTMLGlob("resources/*")
	router.GET("/", renderHTML)
	router.GET("/version", version)
	router.GET("/ws/echo", func(c *gin.Context) {
		echo(c.Writer, c.Request)
	})
	router.Run(":8080")
}

func renderHTML(c *gin.Context) {
	_, err := resources.Asset("resources/websocket.html")
	if err != nil {
		// Asset was not found.
	}
	c.HTML(http.StatusOK, "websocket.html", gin.H{})
}

func version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "1.0.0" })
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		//noinspection GoPlaceholderCount
		fmt.Printf("Failed to set websocket upgrade: %+v \n", err)
		return
	}
	defer
	func() {
		delete(connections, conn)
		conn.Close()
	}()
	connections[conn] = true
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for key, _ := range connections {
			key.WriteMessage(t, msg)
		}
	}
}

func getLogger() *logger.Logger {
	log, err := logger.New("LOG", 1, os.Stdout)
	if err != nil {
		panic(err) // TODO Check for error
	}
	return log
}
