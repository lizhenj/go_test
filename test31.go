package main

import (
	"fmt"
	"time"
)

func main31() {
	c := make(chan int, 50)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
