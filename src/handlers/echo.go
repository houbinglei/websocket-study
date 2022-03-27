package handlers

import (
	"awesomeProject1/src/core"
	"net/http"
)

func Echo(w http.ResponseWriter, req *http.Request) {
	client, _ := core.Upgrader.Upgrade(w, req, nil) //client 是客户端连接对象
	core.ClientMap.Store(client)
}
