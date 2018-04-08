//add registration for clients and maybe store details in a database?
//create frontend of application
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: CheckOrigin} //minimise global variable use
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	messageType int
	message     []byte
}

func CheckOrigin(_ *http.Request) bool {
	return true
}

//maybe put websocket stuff into different files
func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()   //close socket once everything is done
	clients[conn] = true //client is connected

	for {

		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break
		}

		message := Message{messageType: messageType, message: data}
		broadcast <- message //sends message to broadcast channel
	}
}

func handleMessage() {
	for {
		message := <-broadcast        //takes next message from broadcast channel
		for client := range clients { //writes message to each client
			err := client.WriteMessage(message.messageType, message.message)
			if err != nil {
				log.Printf(err.Error())
				client.Close()
				delete(clients, client) //deletes client if they can't receive a message
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessage()
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
