package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func init() {
 	fmt.Println("init() .... ")
	fmt.Printf("%s / %s  Version: %s, CPU NUMBER = %d\n", runtime.GOARCH, runtime.GOOS, runtime.Version(), runtime.NumCPU() )
}

// 原子加，减
func prt_atomic() {
	fmt.Println("prt_atomic() .... ")

	var a int32
	a += 10
	sum := atomic.AddInt32(&a, 10)
	fmt.Println(a == 20) // true
	fmt.Println("sum :" , sum ) // true

	var b uint32
	b += 20
	atomic.AddUint32(&b, ^uint32(10-1))
	// 等价于 b -= 10
	// atomic.Adduint32(&b, ^uint32(N-1))
	fmt.Println(b == 10) // true

}


func main() {
	fmt.Println("Hello, Map!")
	prt_atomic()

}

