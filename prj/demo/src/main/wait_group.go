package main

import (
	"fmt"
	"sync"
)
//sync.WaitGroup是一个计数的信号量， 使main函数所在主线程等待两个goroutine
// 执行完成后再结束，
// 否则两个goroutine还在运行时，主线程已经结束。
//sync.WaitGroup使用非常简单，
// 1、使用Add方法设设置计数器为2，每一个goroutine的函数执行完后，调用Done方法减1。
// 2、Wait方法表示如果计数器大于0，就会阻塞，main函数会一直等待2个goroutine完成再结束。
func main() {
	fmt.Println("Hello sync.WaitGroup ")
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 10001; i++ {
			if  i % 2000 == 0 {
				fmt.Printf("func() 1 , This is %d\n", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10001; i++ {
			if  i % 2000 == 0 {
				fmt.Printf("func() 2 , This is %d\n", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10001; i++ {
			if  i % 2000 == 0 {
				fmt.Printf("func() 3 , This is %d\n", i)
			}
		}
	}()
	wg.Wait()
}
