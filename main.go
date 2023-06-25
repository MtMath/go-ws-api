package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

func main() {
	go MonitorClientConnections()

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ws", WsHandler)


	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Handle("/", http.FileServer(http.Dir("public")))
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	CheckError(err)

	AddClient(conn)

	defer CleanupClients()

	go HandleClientMessages(conn)
}

func AddClient(client *websocket.Conn) {
	clients = append(clients, *client)
	log.Printf("Added client %s\n", client.RemoteAddr())
}

func BroadcastMessageToAll(msg []byte) {
	for _, conn := range clients {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func BroadcastMessageToClient(msg []byte, client *websocket.Conn) {
	client.WriteMessage(websocket.TextMessage, msg)
}

func BroadcastMessageToOthers(msg []byte, client *websocket.Conn) {
	for _, conn := range clients {
		if conn.RemoteAddr() != client.RemoteAddr() {
			err := conn.WriteMessage(websocket.TextMessage, msg)
			CheckError(err)
		}
	}
}

func HandleClientMessages(conn *websocket.Conn) {
	for {
		if clients != nil {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error during message reading:", err)
				break
			}

			log.Printf("%s, Message Received: %s\n", conn.RemoteAddr(), string(msg))

			BroadcastMessageToOthers(msg, conn)
		}
	}
}

func CleanupClients() {
	for i := len(clients) - 1; i >= 0; i-- {
		if clients[i].WriteMessage(websocket.PingMessage, nil) != nil {
			clients[i].Close()
			clients = append(clients[:i], clients[i+1:]...)
		}
	}
}

func MonitorClientConnections() {
	for {
		for i := len(clients) - 1; i >= 0; i-- {
			if clients[i].WriteMessage(websocket.PingMessage, nil) != nil {
				clients[i].Close()
				clients = append(clients[:i], clients[i+1:]...)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
