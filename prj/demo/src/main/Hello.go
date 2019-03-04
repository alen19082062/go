package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	t0 := time.Now()      // 起点时间
	fmt.Println("init() .... ")
	fmt.Println("CPU NUMBERS: " , runtime.NumCPU())
	fmt.Println("GOOS : " , runtime.GOOS)
	fmt.Println("GOARCH : " , runtime.GOARCH)
	fmt.Println("GO Version : " , runtime.Version())
	fmt.Println("GO NumGoroutine : " , runtime.NumGoroutine())

	var j int
	for i:=0 ; i<2000000; i++ {
		j = i + 1 ;
		j = j - 1 ;
	}
	endTime := time.Since(t0) // 耗时
	fmt.Println("j = " , j);
	fmt.Println("init() 运行时间 : ", endTime)
}

func main() {
	fmt.Println("Hello, go World!")
}

