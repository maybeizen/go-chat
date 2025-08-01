package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func RunClient() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)

	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer conn.Close()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			fmt.Println(string(msg))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a name > ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Println("Send a message >")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "/exit" {
			break
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", username, text)))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}