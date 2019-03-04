package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(4)  //使用多核
}

// 无论 cpu 数目为 4 或者 1 ，输出结果是不相同的
// 有时 go func() 会先执行
func main() {
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()


	for i := 0; i < 4; i++ {
		fmt.Println("a:", i)

		if i == 1 {
			fmt.Println("GO NumGoroutine : " , runtime.NumGoroutine())
			runtime.Gosched()  //切换任务
		}
	}
	<-exit

}