package main

import (
	"fmt"
	"runtime"
)

func init() {
 	fmt.Println("init() .... ")
	fmt.Printf("[%s / %s] CPU NUMBER = %d  Version: %s,\n", runtime.GOARCH, runtime.GOOS, runtime.NumCPU(), runtime.Version() )
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func f() {
	fmt.Println("f() begin .... ")

	for i := 0; i < 5; i++ {
		// 这里最后执行
		defer fmt.Printf("defer : %d \n", i)
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("f() %d \n", i)
	}

	fmt.Println("f() end .... ")
}

func main() {
	fmt.Println("Hello, go World!")
	f();

	x := min(1, 3, 2, 7)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

