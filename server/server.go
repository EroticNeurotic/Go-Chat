//add registration for clients and maybe store details in a database?
//create frontend of application
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//we want a message to contain the message itself and the sender (instead of username could use a client type)
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{} //minimise global variable use
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

//does this need to be an entire separate func?
func handleFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html") //make an actual file to be used, maybe replace with file server
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
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break
		}
		broadcast <- message //sends message to broadcast channel
	}
}

func handleMessage() {
	for {
		message := <-broadcast        //takes next message from broadcast channel
		for client := range clients { //writes message to each client
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf(err.Error())
				client.Close()
				delete(clients, client) //deletes client if they can't receive a message
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handleFile)
	http.HandleFunc("/ws", handleConnections)
	go handleMessage()
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
