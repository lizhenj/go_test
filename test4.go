package main

import (
	"log"
	"runtime"
)

func main4() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	where()
	where()
}
