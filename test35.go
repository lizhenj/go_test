package main

import (
	"fmt"
	"os"
	"time"
)

func main35() {
	term := 25
	i := 0
	c := make(chan int)
	end := make(chan bool)
	start := time.Now()

	go fibnterms(term, c, end)
	for {
		select {
		case res := <-c:
			fmt.Printf("fibonacci(%d) is: %d\n", i, res)
			i++
		case <-end:
			delta := time.Since(start)
			fmt.Printf("longCalculation took this amount of time: %s\n", delta)
			os.Exit(0)
		}
	}
}

func fibnterms(term int, c chan int, e chan bool) {
	for i := 0; i <= term; i++ {
		c <- fibonacci(i)
	}
	e <- true
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
