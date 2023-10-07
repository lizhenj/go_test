package main

import (
	"fmt"
	"os"
)

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
}

// fatal error: all goroutines are asleep - deadlock
func main34_1() {
	var ok = true
	ch := make(chan int)

	go tel(ch)
	for ok {
		i := <-ch
		fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	}
}

// fatal error: all goroutines are asleep - deadlock
func main34_2() {
	var ok = true
	var i int
	ch := make(chan int)

	go tel(ch)
	for ok {
		if i, ok = <-ch; ok {
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		} else {
			fmt.Println(ok)
		}
	}
}

func tel1(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func main34() {
	var ok = true
	ch := make(chan int)
	quit := make(chan bool)

	go tel1(ch, quit)
	for ok {
		select {
		case i := <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			os.Exit(0)
		}
	}
}
