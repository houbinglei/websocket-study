package core

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// 全局对象， upgrader 对象
var Upgrader websocket.Upgrader

func init() {
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
