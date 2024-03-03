package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
	"gend.com/server/types"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("Received message: %s", message)

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := types.Articles{
		types.Article{Title: "Title 1", Desc: "Description 1", Content: "Content 1"},
		types.Article{Title: "Title 2", Desc: "Description 2", Content: "Content 2"},
	}
	fmt.Println("Hit all articles")
	json.NewEncoder(w).Encode(articles)
}

func handleRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/websocket", websocketHandler)
}

func main() {
	handleRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
