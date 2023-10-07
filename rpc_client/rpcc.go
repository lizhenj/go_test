package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func mainr() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
		return
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(reply)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
		return
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(reply)
}
