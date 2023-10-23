package main

import "fmt"

func main() {
	kk := []int{1, 3, 5, 7, 9, 4, 2}
	maopao(kk)

	fmt.Println("快速排序前 kk：", kk)
	//quick(kk, 0, len(kk)-1)
	insertSort2(kk)
	fmt.Println("快速排序后 kk：", kk)
}

func maopao(arr []int) {
	arr1 := make([]int, len(arr))
	copy(arr1, arr)

	fmt.Println("排序前 arr：", arr1)
	for i := 0; i < len(arr1)-1; i++ {
		for j := 0; j < len(arr1)-1-i; j++ {
			if arr1[j] > arr1[j+1] {
				arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
			}
		}
	}
	fmt.Println("排序后 arr：", arr1)
}

func quickSort(sequence []int, low int, high int) {
	if high <= low {
		return
	}
	j := partition(sequence, low, high)
	quickSort(sequence, low, j-1)
	quickSort(sequence, j+1, high)
}

// 进行快速排序中的一轮排序
func partition(sequence []int, low int, high int) int {
	i, j := low+1, high
	for {
		// 把头元素作为基准值 pivot
		for sequence[i] < sequence[low] {
			// i 坐标从前往后访问序列，如果位置上的值大于基准值，停下来。
			// 准备和 j 坐标访问到的小于基准值的值交换位置
			i++
			if i >= high {
				break
			}
		}
		for sequence[j] > sequence[low] {
			// j 坐标从后往前访问序列，如果位置上的值小于基准值，停下来。
			// 和 i 坐标指向的大于基准值的值交换位置
			j--
			if j <= low {
				break
			}
		}
		if i >= j {
			break
		}
		sequence[i], sequence[j] = sequence[j], sequence[i]
	}
	sequence[low], sequence[j] = sequence[j], sequence[low]

	return j
}

func quick(arr []int, low, high int) {
	if low >= high {
		return
	}

	//获取基准中间位
	j := calcjizhun(arr, low, high)
	quick(arr, low, j-1)
	quick(arr, j+1, high)
}

func calcjizhun(arr []int, low, high int) int {
	i, j := low+1, high
	for {
		for arr[i] < arr[low] {
			i++
			if i >= high {
				break
			}
		}

		for arr[j] > arr[low] {
			j--
			if j <= low {
				break
			}
		}

		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}

func insertSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		tmp := numbers[i]
		// 从待排序序列开始比较,找到比其小的数
		j := i
		for j > 0 && tmp < numbers[j-1] {
			numbers[j] = numbers[j-1]
			j--
		}
		// 存在比其小的数插入
		if j != i {
			numbers[j] = tmp
		}
	}
}

func insertSort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for j > 0 && tmp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		if j != i {
			arr[j] = tmp
		}
	}
}
