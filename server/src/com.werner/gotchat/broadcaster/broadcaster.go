package broadcaster

import (
  "github.com/gorilla/websocket"
)

type Broadcaster struct {
  connections map[*websocket.Conn]bool
  broadcastChannel *chan Message
}

func NewBroadCaster(broadcastChannel *chan Message) *Broadcaster{
  broadcaster := new(Broadcaster)
  broadcaster.connections = make(map[*websocket.Conn]bool)
  broadcaster.broadcastChannel = broadcastChannel
  return broadcaster
}

type Messenger interface {
  GetMessage() string
}

type Message struct {
  connection *websocket.Conn
  messageType MessageType
  message string
}

func (m *Message) GetMessage() string {
  switch m.messageType {
  case SEND:
    return m.message
  default:
    return nil
  }
}

type MessageType int

const (
  SUBSCRIBE MessageType = iota
  UNSUBSCRIBE
  SEND
)