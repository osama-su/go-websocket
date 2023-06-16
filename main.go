package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// Upgrader - take http connection and upgrade it to a websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/", serveHome)
	fmt.Printf("Server started on port 8080\n")
	fmt.Printf("Visit http://localhost:8080/ to chat\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func serveHome(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "index.html")
}

func echoHandler(writer http.ResponseWriter, request *http.Request) {
	conn, _ := upgrader.Upgrade(writer, request, nil) // error ignored. TODO: handle error
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(message))

		if err = conn.WriteMessage(messageType, message); err != nil {
			return
		}
	}
}
