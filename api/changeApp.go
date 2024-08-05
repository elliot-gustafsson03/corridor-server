package api

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  256,
	WriteBufferSize: 256,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var conn *websocket.Conn = nil

func AddChangeAppEndpoints() {
	http.HandleFunc("/ws/connect", createConnection)
	http.HandleFunc("/api/change_app", changeApp)
}

func createConnection(w http.ResponseWriter, r *http.Request) {
	if conn != nil {
		conn.Close()
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	conn = c
}

func changeApp(w http.ResponseWriter, r *http.Request) {
	if conn != nil {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		conn.WriteMessage(websocket.TextMessage, body)
		io.WriteString(w, "1")
	} else {
		io.WriteString(w, "0")
	}
}
