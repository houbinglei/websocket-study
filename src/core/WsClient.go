package core

import (
	"github.com/gorilla/websocket"
	"time"
)

type WsClient struct {
	conn *websocket.Conn
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn}
}
func (this *WsClient) Ping(waittime time.Duration) {

	for {
		time.Sleep(waittime)
		err := this.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			ClientMap.Remove(this.conn)
			return
		}

	}
}
