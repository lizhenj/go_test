package main

import "fmt"

func main3() {
	g := reFecu()

	for i := 0; i < 30; i++ {
		fmt.Println(g())
	}
}

func reFecu() func() int {
	one, two := 1, 1
	return func() int {
		one, two = two, one+two
		return two
	}
}
