package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct{}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{}
}

func (h *WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// 初回メッセージ
	err = ws.WriteMessage(websocket.TextMessage, []byte(`Server (gorilla): Hello, Client!`))
	if err != nil {
		log.Println("WriteMessage:", err)
		return
	}

	for {
		//mt=message tupe
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("ReadMessage:", err)
			break
		}
		err = ws.WriteMessage(mt, []byte(fmt.Sprintf("Server (gorilla): '%s' received.", message)))
		if err != nil {
			log.Println("WirteMessage:", err)
			break
		}
	}

}
