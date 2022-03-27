package core

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

var ClientMap *ClientMapData

type ClientMapData struct {
	data sync.Map
}

func init() {
	ClientMap = &ClientMapData{}
}

func (this *ClientMapData) Store(conn *websocket.Conn) {
	wsClient := NewWsClient(conn)
	this.data.Store(conn.RemoteAddr().String(), wsClient)
	go wsClient.Ping(time.Second * 1)
}

//向所有客户端 发送消息
func (this *ClientMapData) SendAll(msg string) {
	this.data.Range(func(key, value interface{}) bool {
		c := value.(*WsClient).conn
		err := c.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			this.Remove(c)
			log.Println(err)
		}
		return true
	})
}
func (this *ClientMapData) Remove(conn *websocket.Conn) {
	this.data.Delete(conn.RemoteAddr().String())
}
