package main

import (
	"container/list"
	"fmt"
	"runtime"
)

func init() {
 	fmt.Println("init() .... ")
	fmt.Printf("%s / %s  Version: %s, CPU NUMBER = %d\n", runtime.GOARCH, runtime.GOOS, runtime.Version(), runtime.NumCPU() )
}


func list_prt() {
	fmt.Println("list_prt() begin .... ")

	// var l list.List

	l := list.New()
	l.PushBack("fist")
	l.PushFront(67)

	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(1)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)

	fmt.Println("print list ---> " ) ;
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("list_prt() end .... ")
}


func main() {
	fmt.Println("Hello, list !")

	list_prt()
}

