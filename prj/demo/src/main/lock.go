package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	//声明
	var mutex sync.Mutex
	fmt.Println("Lock the mutex. (Main G0)")
	//加锁mutex
	mutex.Lock()

	fmt.Println("The lock is locked.(Main G0)")
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (Coroutine %d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (Goroutine %d)\n", i)
		}(i)
	}
	//休息一会,等待打印结果
	// time.Sleep( 2 * time.Second)
	fmt.Println("Unlock the lock. (Main G0)")
	//解锁mutex
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (Main G0)")

	time.Sleep( time.Second)
	//解锁mutex
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (Main G0)")

	// 如果不休眠，会报错, fatal error: sync: unlock of unlocked mutex
	time.Sleep(  time.Second)
	//解锁mutex
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (Main G0)")

	//休息一会,等待打印结果
	time.Sleep(time.Second)
}

