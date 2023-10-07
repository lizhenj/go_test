package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"unsafe"
)

func main24() {
	//CopyFile("dst.txt", "test.txt")
	fmt.Println("Copy done!")
	//r := gin.Default()
	//
	//r.GET("/hello", func(c *gin.Context) {
	//	c.JSON(200, map[string]string{
	//		"hhh": "hhhh",
	//	})
	//})
	//r.Run(":9090")
	a := []int{1, 2, 3}
	c := a[:0]
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", c)
	fmt.Println(copy(a, a[1:]))
	fmt.Println(a)
	a = a[:copy(a, a[1:])]
	fmt.Println(a)
	copy(a, a[2:])
	fmt.Println(a)
	s := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Println(s)
	fmt.Println(f1(5))
	fmt.Println(g())
}

func f1(x int) *int {
	return &x
}

func g() int {
	x := new(int)
	return *x
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
