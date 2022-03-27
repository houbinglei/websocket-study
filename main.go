package main

import (
	"awesomeProject1/src/core"
	"awesomeProject1/src/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", handlers.Echo)

	http.HandleFunc("/sendall", func(w http.ResponseWriter, req *http.Request) {
		msg := req.URL.Query().Get("msg")
		core.ClientMap.SendAll(msg)
		w.Write([]byte("OK"))

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
