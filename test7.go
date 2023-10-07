package main

import "fmt"

func main7() {
	sla := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ", sla)
	bubbleSort(sla)
	fmt.Println("after sort: ", sla)
}

// 冒泡排序
func bubbleSort(sl []int) {
	for i := 1; i < len(sl); i++ {
		for j := 0; j < len(sl)-i; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
}
