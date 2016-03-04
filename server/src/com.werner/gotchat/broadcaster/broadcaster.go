package broadcaster

import (
  "os"
  "log"
  "github.com/gorilla/websocket"
)

var LOGGER *log.Logger = log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)

type Broadcaster struct {
  connections map[*websocket.Conn]bool
  broadcastChannel chan *MessageWrapper
}

func NewBroadCaster(broadcastChannel chan *MessageWrapper) *Broadcaster{
  broadcaster := new(Broadcaster)
  broadcaster.connections = make(map[*websocket.Conn]bool)
  broadcaster.broadcastChannel = broadcastChannel
  return broadcaster
}

func (b *Broadcaster) HandleMessages(){
  for true {
    msgWrapper := <- b.broadcastChannel
    //LOGGER.Println(msgWrapper.GetMessage())
    switch msgWrapper.operation {
    case SUBSCRIBE:
      b.connections[msgWrapper.connection] = true
      break
    case UNSUBSCRIBE:
      delete(b.connections, msgWrapper.connection)
      break
    case SEND:
      for key, _ := range b.connections {
	key.WriteMessage(msgWrapper.messageType, []byte(msgWrapper.message))
      }
      break
    default:
      continue
    }
  }
  return
}

//type Messenger interface {
//  GetMessage() string
//}

type MessageWrapper struct {
  connection *websocket.Conn
  operation  MessageWrapperOperation
  message    string
  messageType int
}

func NewMessageWrapper(connection *websocket.Conn, msg string, msgType int, operation MessageWrapperOperation) *MessageWrapper {
  messageWrapper := new(MessageWrapper)
  messageWrapper.connection = connection
  messageWrapper.operation = operation
  messageWrapper.message = msg
  messageWrapper.messageType = msgType
  return messageWrapper
}

func (m *MessageWrapper) GetMessage() string {
  switch m.operation {
  case SEND:
    return m.message
  default:
    return ""
  }
}

type MessageWrapperOperation int

const (
  SUBSCRIBE MessageWrapperOperation = iota
  UNSUBSCRIBE
  SEND
)