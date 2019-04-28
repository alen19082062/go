package main

import (
	"fmt"
	"runtime"
)

func init() {
	// 用来设置可以并行计算的CPU核数最大值，并返回之前的值。
	runtime.GOMAXPROCS(4)
}

// 无论 cpu 数目为 4 或者 1 ，输出结果是不相同的
// 有时 go func() 会先执行
func main() {
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("Goroutine b")
		}()
	}()


	for i := 0; i < 4; i++ {
		fmt.Println("Main() a:", i)

		if i == 1 {
			fmt.Println("Goroutine Numbers : " , runtime.NumGoroutine())
			// 用于让出CPU时间片
			runtime.Gosched()
		}
	}
	<-exit

}