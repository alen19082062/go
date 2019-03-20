package main

import (
	"fmt"
	"sync"
)

func producer3(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i:=0 ; i< n; i++  {
			out <- i
		}
	}()
	return out
}

func square3(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
			// simulate
			// time.Sleep(2 * time.Second)
		}
	}()
	return out
}

// fan in
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int,100)

	var wg sync.WaitGroup
	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}
	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := producer3(2000000)

	// FAN-OUT
	c1 := square3(in)
	c2 := square3(in)
	c3 := square3(in)

	// consumer
	var count int
	count = 0
	for ret := range merge(c1, c2, c3) {
		count++

		if count % 100 == 0{
			fmt.Printf("count = %d , %5d", count, ret)
		}

		// 换行 ，3个一行
		if count % 100 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()

}