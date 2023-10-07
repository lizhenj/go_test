package main

import (
	"container/list"
	"fmt"
)

func main8() {
	lst := list.New()
	lst.PushBack(100)
	lst.PushBack(101)
	lst.PushBack(102)
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		fmt.Printf("%v", e)
	}
}
