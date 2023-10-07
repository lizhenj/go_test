package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
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
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	dialer := &websocket.Dialer{}
	header := http.Header{
		"name": []string{"Tome", "Jim"},
	}
	conn, resp, err := dialer.Dial("ws://127.0.0.1:3434/add", header)
	CheckError(err)
	for key, values := range resp.Header {
		fmt.Printf("%s:%v\n", key, values)
	}
	defer conn.Close()

	for {
		request := Request{A: 3, B: 9}
		err = conn.WriteJSON(request)
		CheckError(err)

		var response Response
		err = conn.ReadJSON(&response)
		fmt.Printf("response sum=%d\n", response.Sum)
		time.Sleep(time.Second * 10)
	}
}
