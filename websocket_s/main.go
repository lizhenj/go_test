package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type (
	Request struct {
		A int
		B int
	}
	Response struct {
		Sum int
	}
	WsServer struct {
		listener net.Listener
		addr     string
		upgrade  *websocket.Upgrader
	}
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewWsServer(port int) *WsServer {
	ws := new(WsServer)
	ws.addr = "0.0.0.0:" + strconv.Itoa(port)
	ws.upgrade = &websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {

		},
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	return ws
}

func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add" {
		httpCode := http.StatusInternalServerError
		phrase := http.StatusText(httpCode)
		http.Error(w, phrase, httpCode)
		return
	}
	for key, values := range r.Header {
		fmt.Printf("%s:%v\n", key, values)
	}
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("upgrade from http to websocket failed: %v\n", err)
		return
	}
	defer conn.Close()
	_ = conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	for {
		var request Request
		err = conn.ReadJSON(&request)
		if err != nil {
			fmt.Printf("Mage read error: %v\n", err)
			break
		}
		fmt.Printf("receive request a=%d b=%d\n", request.A, request.B)
		sum := request.A + request.B
		response := Response{
			Sum: sum,
		}
		err = conn.WriteJSON(&response)
		CheckError(err)
		time.Sleep(time.Second * 8)
	}
}

func main() {
	ws := NewWsServer(3434)
	listener, err := net.Listen("tcp", ws.addr)
	CheckError(err)
	ws.listener = listener
	err = http.Serve(listener, ws)
	CheckError(err)
}
