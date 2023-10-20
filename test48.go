package main

import (
	"fmt"
	"math/big"
	"time"
)

const LIM = 30

var fibs [LIM]*big.Int

func main48() {
	//result := big.NewInt(0)
	var result *big.Int
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci48(i)
		fmt.Printf("数列第%d位：%d\n", i+1, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行完成，所耗时间为：%s\n", delta)
}

func fibonacci48(n int) (res *big.Int) {
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs[n-1], fibs[n-2])
	}
	fibs[n] = res
	return
}
