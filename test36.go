package main

import (
	"fmt"
	"sort"
)

func fibonacci1(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main36_1() {
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

type info struct {
	id    int
	socre int
}

func main36() {
	kk := make([]*info, 0)
	for i := 1; i <= 10; i++ {
		info1 := &info{id: i}
		kk = append(kk, info1)
	}
	kk[0].socre = 10
	kk[1].socre = 10
	kk[2].socre = 15
	kk[3].socre = 15
	kk[4].socre = 20
	kk[5].socre = 20
	kk[6].socre = 40
	kk[7].socre = 40
	kk[8].socre = 10
	kk[9].socre = 10
	printArr(kk)
	sort.Slice(kk, func(i, j int) bool {
		return kk[i].socre > kk[j].socre
	})
	printArr(kk)
}

func printArr(arr []*info) {
	for _, info1 := range arr {
		fmt.Println(info1.id, info1.socre)
	}
	fmt.Println("------------------")
}
