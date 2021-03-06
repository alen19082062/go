package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)

	// 从 第三开始
	go sum(s[:len(s)/2], ch)
	// 到 第三结束
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch // receive from c
	fmt.Println(x, y, x+y)
}

