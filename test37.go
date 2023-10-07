package main

import "fmt"

func main37() {
	c, quit := make(chan int), make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 1
	}()
	fibonacci2(c, quit)
}

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("end")
			return
		}
	}
}
