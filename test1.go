package main

import "fmt"

func main1() {
	for i := uint64(0); i <= uint64(10); i++ {
		fmt.Printf("the %v result :%v\n", i, calcJie(i))
	}
}

func calcJie(n uint64) (fac uint64) {
	fac = 1

	if n >= 1 {
		fac = n * calcJie(n-1)
		return fac
	}
	return fac
}
