package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	t0 := time.Now()      // 起点时间
	fmt.Println("init() .... ")
	fmt.Printf("[%s / %s] CPU NUMBER = %d  Version: %s,\n", runtime.GOARCH, runtime.GOOS, runtime.NumCPU(), runtime.Version() )
	// ？？？ 此处有问题，怎么只有一个协程
	// runtime.ReadMemStats();
	// runtime.StartTrace()
	fmt.Println("GO NumGoroutine : " , runtime.NumGoroutine())

	var j int
	for i:=0 ; i<1000000; i++ {
		j = i + 1 ;
		j = j - 1 ;
	}
	endTime := time.Since(t0) // 耗时
	fmt.Println("j = " , j);
	fmt.Println("init() 运行时间 : ", endTime)
}

func main() {
	fmt.Println("Hello, go World!")

	fmt.Printf("[%s / %s] CPU NUMBER = %d  Version: %s,\n", runtime.GOARCH, runtime.GOOS, runtime.NumCPU(), runtime.Version() )
	fmt.Println("GO NumGoroutine : " , runtime.NumGoroutine())

}

