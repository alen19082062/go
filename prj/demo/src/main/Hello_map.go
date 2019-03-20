package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

func init() {
 	fmt.Println("init() .... ")
	fmt.Printf("%s / %s  Version: %s, CPU NUMBER = %d\n", runtime.GOARCH, runtime.GOOS, runtime.Version(), runtime.NumCPU() )
}

func map_prt() {
	var m1 map[string]string
	var m2 map[string]string = map[string]string{}      // m2 := map[string]string{}
	var m3 map[string]string = make(map[string]string, 10)  // m3 := make(map[string]string)

	//m1["1"] = "1"   // panic: assignment to entry in nil map
	m2["2"] = "2"
	m3["3"] = "3"

	for key, value := range m1 { fmt.Println("Key:", key, "Value:", value) }
	for key, value := range m2 { fmt.Println("Key:", key, "Value:", value) }
	for key, value := range m3 { fmt.Println("Key:", key, "Value:", value) }

	s1 := m1["1"]
	s2 := m2["2"]
	s3 := m3["3"]

	fmt.Printf("val=%s,%s,%s\n", s1, s2, s3)
	fmt.Printf("len=%d,%d,%d\n", len(m1), len(m2), len(m3))
	fmt.Printf("size=%d,%d,%d\n", unsafe.Sizeof(m1), unsafe.Sizeof(m2), unsafe.Sizeof(m3))

	//PrintMemory(unsafe.Pointer(&m1), 8)
	//PrintMemory(unsafe.Pointer(&m2), 8)
	//PrintMemory(unsafe.Pointer(&m3), 8)
}

func sync_map(){
	fmt.Println("sync_map() .... ")

	//开箱即用
	var sm sync.Map
	//store 方法,添加元素
	sm.Store(1,"a")
	sm.Store(4,"d")

	//Load 方法，获得value
	if v,ok:=sm.Load(1);ok{
		fmt.Println("Found in map , value : " ,v)
	} else {
		fmt.Println("Not found in map " )
	}

	v,ok := sm.LoadOrStore("1","aaa")
	fmt.Println("LoadOrStore(1,aaa) , " ,ok,v) //false aaa

	// 若key已存在，则返回true和key对应的value，不会修改原来的value
	v,ok = sm.LoadOrStore(1,"aaa")
	fmt.Println("LoadOrStore(1,aaa) , " ,ok,v) //false aaa

	// LoadOrStore方法，获取或者保存
	// 参数是一对key：value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；
	// 不存在则store，返回该value 和false
	if vv,ok:=sm.LoadOrStore(1,"c");ok{
		fmt.Println("Found in map , value : " ,vv)
	}else {
		fmt.Println("Not found in map " )
	}

	if vv,ok:=sm.LoadOrStore(2,"c");!ok{
		fmt.Println("Found in map , value : " ,vv)
	}else {
		fmt.Println("Not found in map " )
	}

	// 遍历该map，参数是个函数，该函数参的两个参数是遍历获得的key和value，
	// 返回一个bool值，当返回false时，遍历立刻结束。
	sm.Range(func(k,v interface{})bool{
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
		return true
	})

}

func main() {
	fmt.Println("Hello, Map!")
	map_prt()

	sync_map()
}

