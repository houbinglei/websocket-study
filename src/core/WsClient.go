package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type WsClient struct {
	conn       *websocket.Conn
	NormalChan chan *WsMsgStruct
	ErrChan    chan int
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn, NormalChan: make(chan *WsMsgStruct), ErrChan: make(chan int)}
}

func (this *WsClient) Ping(waittime time.Duration) {
	for {
		time.Sleep(waittime)
		err := this.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		this.conn.ReadMessage()
		if err != nil {
			fmt.Println("xxx")
			ClientMap.Remove(this.conn)
			return
		}
	}
}

func (this *WsClient) ReadLoop() {
	for {
		t, data, err := this.conn.ReadMessage()
		if err != nil {
			fmt.Println("readloop err is ", err)
			this.conn.Close()
			ClientMap.Remove(this.conn)
			this.ErrChan <- 1
			break
		}
		fmt.Println(777)
		fmt.Println("before chan is :", NewWsMsgStruct(t, data))
		this.NormalChan <- NewWsMsgStruct(t, data)
	}
}

func (this *WsClient) handlerLoop() {
loop:
	for {
		select {
		case <-this.ErrChan:
			log.Println("已经关闭")
			break loop
		case data := <-this.NormalChan:
			fmt.Println(string(data.MessageData))
		}
	}
}
