package main

import "fmt"

func main52() {
	a := "ababaeabacaaaabcqadzcwweasdcaswqeiacawwac"
	b := "acawwac"
	pos := KMP2(b, a)
	fmt.Println(pos)
	pos = KMP(b, a)
	fmt.Println(pos)
}

//KMP算法
func KMP(needle string, str string) int {
	next := getNext(needle)
	fmt.Println(next)
	j := 0
	for i := 0; i < len(str); i++ {
		for j > 0 && str[i] != needle[j] {
			j = next[j-1] + 1
		}
		if str[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func getNext(needle string) []int {
	var next = make([]int, len(needle))
	//fmt.Println(next)
	next[0] = -1
	k := -1
	for i := 1; i < len(needle); i++ {
		for k != -1 && needle[k+1] != needle[i] {
			k = next[k]
		}
		if needle[k+1] == needle[i] {
			k++
		}
		next[i] = k
	}
	return next
}

func KMP2(needle, str string) int {
	next := build_next(needle)
	fmt.Println(next)

	j := 0
	for i := 0; i < len(str); i++ {
		for j > 0 && str[i] != needle[j] {
			j = next[j-1]
		}
		if str[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}

	return -1
}

func build_next(needle string) []int {
	var next = make([]int, len(needle))
	next[0] = 0
	k := 0
	for i := 1; i < len(needle); i++ {
		for k != 0 && needle[k] != needle[i] {
			k = next[k-1]
		}
		if needle[k] == needle[i] {
			k++
		}
		next[i] = k
	}
	return next
}
