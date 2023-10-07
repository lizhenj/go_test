package main

import (
	"fmt"
	"sync"
	"time"
)

func main46_1() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	for v := range ch {
		fmt.Printf("%v", v)
	}
}

func worker(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancel:
			break
		}
	}
}

func main46_2() {
	cancel := make(chan bool)
	go worker(cancel)

	time.Sleep(time.Second)
	cancel <- true
}

func worker1(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancel:
			return
		}
	}
}

func main46() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker1(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()

}

func main46_3() {
	var kk interface{}
	ww := 10
	kk = &ww
	switch v := kk.(type) {
	case int:
		fmt.Println(v)
	case *int:
		*v = 18
		fmt.Println(*v)
	}
	fmt.Println(kk)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

}
