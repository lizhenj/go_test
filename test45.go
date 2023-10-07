package main

import (
	"context"
	"fmt"
	"sync"
)

// 素数筛-调用GenerateNatural生成从2开始的自然数序列。
// 开始一个100次的迭代循环，每次循环迭代的开始，管道第一个元素必然是素数（基于此）
func main45() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	ch := GenerateNatural(ctx, &wg) //自然数序列
	wg.Add(1)
	for i := 0; i < 100; i++ {
		prime := <-ch //新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		wg.Add(1)
		ch = PrimeFilter(ctx, ch, prime, &wg) //基于新素数构造的过滤器
	}
	cancel()
	wg.Wait()
}

// 返回生成自然数序列的管道：2，3，4.。。
func GenerateNatural(ctx context.Context, wg *sync.WaitGroup) chan int {
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func PrimeFilter(ctx context.Context, in <-chan int, prime int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()
		defer close(out)
		//for {
		//select {
		//case <-ctx.Done():
		//	return
		//case i := <-in:
		//	if i%prime != 0 {
		//		out <- i
		//	}
		//}
		//}
		for i := range in {
			if i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}
