package main

import (
  "net/http"
  "os"
  "log"
  "sync"
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
)

var connections map[*websocket.Conn]bool;
// Used to sync access to map connections
var mutex = &sync.Mutex{}

var LOGGER *log.Logger;

func main() {
  LOGGER = log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)
  connections = make(map[*websocket.Conn]bool)
  router := gin.Default()
  router.GET("/version", version)
  router.GET("/ws/echo", func(c *gin.Context) {
    echo(c.Writer, c.Request)
  })
  router.Run(":8080")
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
    LOGGER.Printf("Failed to set websocket upgrade: %+v \n", err)
    return
  }
  defer
  func() {
    mutex.Lock()
    delete(connections, conn)
    mutex.Unlock()
    conn.Close()
  }()
  mutex.Lock()
  connections[conn] = true
  mutex.Unlock()
  for {
    t, msg, err := conn.ReadMessage()
    if err != nil {
      break
    }
    mutex.Lock()
    for key, _ := range connections {
      key.WriteMessage(t, msg)
    }
    mutex.Unlock()
  }
}