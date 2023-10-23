package main

import (
	"fmt"
	"time"
)

var dog = make(chan struct{})
var cat50 = make(chan struct{})
var fish = make(chan struct{})

func Dog() {
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
}

func Cat() {
	<-dog
	fmt.Println("cat")
	cat50 <- struct{}{}
}

func Fish() {
	<-cat50
	fmt.Println("fish")
	fish <- struct{}{}
}

func main50() {
	for i := 0; i < 100; i++ {
		go Dog()
		go Cat()
		go Fish()
	}
	fish <- struct{}{}

	time.Sleep(10 * time.Second)
}
