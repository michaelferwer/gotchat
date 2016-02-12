package broadcaster

import (
  "github.com/gorilla/websocket"
)

type Broadcaster struct {
  connections map[*websocket.Conn]bool
  broadcastChannel *chan string
}

func NewBroadCaster() *Broadcaster{
  return new(Broadcaster)
}

type Message struct {
  connection *websocket.Conn

}