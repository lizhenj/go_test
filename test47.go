package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	data := []byte("hello world!")
//
//	read := bytes.NewReader(data)
//	kk := make([]byte, read.Len())
//
//	n, err := read.Read(kk)
//	fmt.Println(n, err)
//	fmt.Println(string(kk))
//
//	read2 := bytes.NewBuffer(data)
//	fmt.Println(read2.String())
//	fmt.Fprintf(read2, "HHHHH")
//	fmt.Println(read2.String())
//
//	read2.WriteTo(os.Stdout)
//}

// return先defer后 由于使用了命名返回值，导致的r在defer中加一，最终返回一
func a() (r int) {
	defer func() {
		r++
	}()
	return 0
}

// return处决定了r=5，defer对r不影响
func b() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// 参数传递即为值传递，故defer中的r为局部变量  //若传指针则会加5
func c() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main47() {
	fmt.Println("a = ", a())
	fmt.Println("b = ", b())
	fmt.Println("c = ", c())
	calc(func() { fmt.Println("HHHHHH") })
	calc("hhhh")
}

func calc(cbs interface{}) {
	cb := reflect.ValueOf(cbs)
	switch cb.Kind() {
	case reflect.Func:
		kk := []reflect.Value{}
		cb.Call(kk)

		aa := cb.Interface().(func())
		aa()
	default:
		fmt.Println(cb.String())
	}
	fmt.Printf("%v\n", cb.Kind())

}
