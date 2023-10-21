package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections []*websocket.Conn

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	connections = append(connections, conn)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Received message: %s\n", p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}
	}
}

//func send(w http.ResponseWriter, r *http.Request) {
//	for i := 0; i < len(connections); i++ {
//		if err := connections[i].WriteJSON("Test"); err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//}

func main() {
	http.HandleFunc("/ws/", handleConnection)
	//http.HandleFunc("/test-send", send)

	port := ":8090"
	fmt.Printf("Server is listening on port %s\n", port)
	http.ListenAndServe(port, nil)
}
