package main

import (
	"fmt"
	"reflect"
)

func main17() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("the Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
