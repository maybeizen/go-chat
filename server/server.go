package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}
	ws, err := upgrader.Upgrade(w, r, nil)
	
	if err != nil {
		fmt.Println("Websocket Upgrade Error:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- string(msg)
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	fmt.Println("Server running at localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}