package main

import (
  "net/http"
  "os"
  "log"
  "com.werner/gotchat/broadcaster"
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
)

var broadcastChannel chan *broadcaster.MessageWrapper

var LOGGER *log.Logger;

func main() {
  LOGGER = log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)
  broadcastChannel = make(chan *broadcaster.MessageWrapper)

  // Make Broadcaster and lunch listening message in goroutine
  broadcaste := broadcaster.NewBroadCaster(broadcastChannel)
  go func() { broadcaste.HandleMessages() }()

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
  defer func() {
    broadcastChannel <- broadcaster.NewMessageWrapper(conn, "", 0, broadcaster.UNSUBSCRIBE)
    conn.Close()
  }()

  broadcastChannel <- broadcaster.NewMessageWrapper(conn, "", 0, broadcaster.SUBSCRIBE)
  for {
    msgType, msg, err := conn.ReadMessage()
    if err != nil {
      break
    }
    broadcastChannel <- broadcaster.NewMessageWrapper(conn, string(msg), msgType, broadcaster.SEND)
  }
}