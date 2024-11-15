package config

import "github.com/gorilla/websocket"

func makeVariable() {
	mapSocket = make(map[string]*websocket.Conn)
	mapSocketEvent = make(map[string]map[string]*websocket.Conn)

	// chanel job
	emailChan = make(chan EmailJob_MessPayload)
}
