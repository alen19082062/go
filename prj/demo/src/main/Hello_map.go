package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// sync.Map没有Len方法，如果想得到当前Map中有效的entries的数量，需要使用Range方法遍历一次
func init() {
 	fmt.Println("init() .... ")
	fmt.Printf("%s / %s  Version: %s, CPU NUMBER = %d\n", runtime.GOARCH, runtime.GOOS, runtime.Version(), runtime.NumCPU() )
}

func map_prt() {
	fmt.Println("map_prt() begin ")

	var m1 map[string]string
	var m2 map[string]string = map[string]string{}      // m2 := map[string]string{}
	var m3 map[string]string = make(map[string]string, 10)  // m3 := make(map[string]string)

	// panic: assignment to entry in nil map
	//m1["1"] = "1"

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

	// 遍历 map m2
	m2["2"] = "222"
	m2["3"] = "3"
	fmt.Println("traversal this Map m2 ===> " )
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	//PrintMemory(unsafe.Pointer(&m1), 8)
	//PrintMemory(unsafe.Pointer(&m2), 8)
	//PrintMemory(unsafe.Pointer(&m3), 8)

	fmt.Println("map_prt() end ")

}

func sync_map(){
	fmt.Println("sync_map() begin ")

	//开箱即用
	var sm sync.Map
	//store 方法,添加元素
	sm.Store(1,"(1)orgin value")
	sm.Store(4,"(4)orgin value")

	//Load 方法，获得value
	if v,ok:=sm.Load(1);ok{
		fmt.Println("Load(1) Found key 1 in map , value : " ,v)
	} else {
		fmt.Println("Load(1) Not found key 1 in map " )
	}

	// LoadOrStore方法，若key已存在，则返回true和key对应的value，不会修改原来的value
	v,ok := sm.LoadOrStore(1,"(1) new value")
	if  ok {
		fmt.Println("LoadOrStore(1,aaa) key 1 Found in map, value : " ,v) //false aaa
	} else {
		fmt.Println("LoadOrStore(1,aaa) key 1 not found " ) //false aaa
	}

	// 若key不存在，保存提供的键值(Store)。
	v,ok = sm.LoadOrStore(13,"(13)new value added ")
	fmt.Println("LoadOrStore(13,aaa) " ,ok,v) //false aaa

	if vv,ok:=sm.LoadOrStore(13,"second new value");ok{
		fmt.Println("key 13 Found in map , value : " ,vv)
	}else {
		fmt.Println("key 13 Not found in map " )
	}

	if vv,ok:=sm.LoadOrStore(2,"(2)added value" );!ok{
		fmt.Println("key 2 Not Found in map, add new entry")
	}else {
		fmt.Println("key 2 found in map , value : " , vv )
	}

	// 遍历该map，参数是个函数，该函数参的两个参数是遍历获得的key和value，
	// 返回一个bool值，当返回false时，遍历立刻结束。
	fmt.Println("traversal this sync.Map ===> " )
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

