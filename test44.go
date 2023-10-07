package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main44() {
	//使用位运算符判断奇偶性
	i := 101
	fmt.Printf("%b\n", i)
	fmt.Printf("%b\n", (i & 101))

	fmt.Println(101 & 1)
	fmt.Println(100 & 1)

	str := "黑人"
	chars := []rune(strings.ToUpper(str))
	fmt.Println(chars)

	kk := map[int]bool{1: true}
	fmt.Println(kk[2])

	kk2 := map[int]int{1: 0}
	fmt.Println(kk2[2])

	file, _ := os.Open("dst.txt")
	read := bufio.NewReader(file)
	for {
		line, _, err := read.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s\n", line)
	}
	//ss := strings.NewReader("123\n456\n789")
	//read := bufio.NewReader(ss)

	//for {
	//	line, err := read.ReadBytes('\n')
	//	if err == io.EOF {
	//		break
	//	}
	//	linelen := len(line) - 1 //去除换行符
	//	fmt.Printf("%s\n", line[0:linelen])
	//}
	//defer file.Close()
}
