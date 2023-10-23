package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main49() {
	a := []string{"a", "b", "c"}
	//1 + 需遍历字符串，计算并开辟新空间存储
	ret := a[0] + a[1] + a[2]

	//2 fmt.Sprintf 采用了接口参数，必须用反射，因此有性能损耗
	ret = fmt.Sprintf("%s%s%s", a[0], a[1], a[2])

	//3 strings.Builder 内部用指针+切片实现，直接把[]byte转为string，避免变量拷贝
	var sb strings.Builder
	sb.WriteString(a[0])
	sb.WriteString(a[1])
	sb.WriteString(a[2])
	ret = sb.String()

	//4 bytes.Buffer byte缓冲器，底层也为一[]byte切片
	buf := new(bytes.Buffer)
	buf.Write([]byte(a[0]))
	buf.Write([]byte(a[1]))
	buf.Write([]byte(a[2]))
	ret = buf.String()

	//5 strings.Join 基于strings.builder来实现 可自定义分隔符
	ret = strings.Join(a, "")
	fmt.Printf("s\n", ret)

	fmt.Println(11635066408315 >> 32)

	q := make(chan int, 2)
	q <- 1
	q <- 2
	select {
	case q <- 3:
		fmt.Println("ok")
	default:
		fmt.Println("wrong")
	}
}
