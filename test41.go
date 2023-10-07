package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 生产者：生成factor整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main41_1() {
	ch := make(chan int, 64)

	go Producer(3, ch) //生成3的倍数序列
	go Producer(5, ch)
	go Consumer(ch) //消费生成的队列

	time.Sleep(time.Second * 10)
}

func main41_2() {
	ch := make(chan int, 64)

	go Producer(3, ch) //生成3的倍数序列
	go Producer(5, ch)
	go Consumer(ch) //消费生成的队列

	//ctrl c退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func main41() {
	ch := make(chan int)
	go func(ch chan int) {
		fmt.Println("HHH")
		<-ch
	}(ch)

	ch <- 1
	fmt.Println("over")
}
