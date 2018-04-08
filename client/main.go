package main

import (
    "log"
    "net/http"

     "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
    Username    string `json:"username"`
    Content     string `json:"content"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    clients[conn] = true
    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        log.Printf("msg: %v", msg)
        if err != nil {
            log.Printf("error: %v", err)
            delete(clients, conn)
            break
        }
        broadcast <- msg
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        log.Println("Message received")
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("error: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("public")))
    http.HandleFunc("/ws", handleConnections)
    go handleMessages()
    log.Println("Server launched on :9999")
    err := http.ListenAndServe(":9999", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
