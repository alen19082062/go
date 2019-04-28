package main

import (
	"fmt"
	"runtime"
	"time"
)

// 内存使用情况
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

type Garbage struct{ a int }

func notify(f *Garbage) {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)

	fmt.Println("notify() Last GC was:", stats.LastGC)

	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Printf("notify() %d Kb\n",m.Alloc/1024)

	go ProduceFinalizedGarbage()
}

func ProduceFinalizedGarbage() {
	x := &Garbage{}
	// 需要在一个对象object被从内存中移除前执行一些特殊操作
	runtime.SetFinalizer(x, notify)
}

func main() {
	fmt.Println("Hello, go World!")

	fmt.Printf("[%s / %s] CPU NUMBER = %d  Version: %s,\n", runtime.GOARCH, runtime.GOOS, runtime.NumCPU(), runtime.Version() )
	fmt.Println("GO NumGoroutine : " , runtime.NumGoroutine())

	go ProduceFinalizedGarbage()

	for {
		go ProduceFinalizedGarbage()

		fmt.Println("GC() begin ...  " )
		runtime.GC()
		fmt.Println("GC() end ...  " )
		time.Sleep(10 * time.Second) // Give GC time to run

	}
}

