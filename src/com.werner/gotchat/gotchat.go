package main

import (
	"net/http"
	"os"
	"fmt"
	"github.com/apsdehal/go-logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	router := gin.Default()
	router.GET("/version", version)
	router.GET("/ws/echo", func(c *gin.Context) {
        echo(c.Writer, c.Request)
    })
	router.Run(":8080")
}

func version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "version": "1.0.0" })
}

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request){
	conn, err := wsupgrader.Upgrade(w, r, nil)
	    if err != nil {
	        fmt.Printf("Failed to set websocket upgrade: %+v \n", err)
	        return
	    }

	    for {
	        t, msg, err := conn.ReadMessage()
	        if err != nil {
	            break
	        }
	        conn.WriteMessage(t, msg)
	    }
}

func getLogger() *logger.Logger {
	log, err := logger.New("LOG", 1, os.Stdout)
	if err != nil {
		panic(err) // TODO Check for error
	}
	return log
}
