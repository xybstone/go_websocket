package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	ConnectionMax = 100 //cs max Connection
)

var (
	cmdLeg = 7           //命令长度
	bufLeg = 1024 * 1024 //字节长度
)

type myServer struct {
}

func (ms myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  bufLeg,
		WriteBufferSize: bufLeg,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	WsServer(conn)
}

func WsServer(ws *websocket.Conn) {
	for {
		_, buf, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		if err != nil {
			ws.WriteMessage(websocket.BinaryMessage, []byte(err.Error()))
			break
		}

		go doServer(ws, buf)

	}
	fmt.Printf("service finish %s\n", time.Now().String())
}

func doServer(ws *websocket.Conn, body []byte) {
	ws.WriteMessage(websocket.BinaryMessage, body)
}

func main() {
	fmt.Printf("Welcome lcsoft xt server!")
	h := new(myServer)

	err := http.ListenAndServe(":7001", h)
	if err != nil {
		fmt.Printf("ListenAndServe: " + err.Error())
	}

}
