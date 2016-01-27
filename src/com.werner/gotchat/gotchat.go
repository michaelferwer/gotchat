package main

import (
	"net/http"
	"os"
	"fmt"
	"github.com/apsdehal/go-logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var connections [10]*websocket.Conn;
var size int = 0;

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
		CheckOrigin: func(r *http.Request) bool { return true },
}

func echo(w http.ResponseWriter, r *http.Request){
	conn, err := wsupgrader.Upgrade(w, r, nil)
	    if err != nil {
	        fmt.Printf("Failed to set websocket upgrade: %+v \n", err)
	        return
	    }
			connections[size] = conn
			size ++
	    for {
	        t, msg, err := conn.ReadMessage()
	        if err != nil {
	            break
	        }
					for i := 0; i < size ; i ++ {
							connections[i].WriteMessage(t, msg)
					}
	    }
			for i := 0; i < size ; i ++ {
				if conn == connections[i] {
					size --
					connections[i] = nil
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
