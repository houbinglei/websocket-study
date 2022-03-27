package core

import (
	"github.com/gorilla/websocket"
	"sync"
)

var Client *ClientMap

type ClientMap struct {
	data sync.Map
}

func init() {
	Client = &ClientMap{}
}

func (this *ClientMap) Store(key string, client *websocket.Conn) *ClientMap {
	this.data.Store(key, client)
	return this
}

func (this *ClientMap) SendAll(str string) {
	this.data.Range(func(key, value any) bool {
		value.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(str))
		return true
	})
}
