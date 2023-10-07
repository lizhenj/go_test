package main

import (
	"fmt"
	"strings"
)

func main2() {
	asciiOnly := func(c rune) rune {
		if c > 127 {
			return ' '
		}
		return c
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))
}
