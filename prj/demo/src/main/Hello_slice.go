package main

import (
	"fmt"
	"runtime"
)

func init() {
 	fmt.Println("init() .... ")
	fmt.Printf("%s / %s  Version: %s, CPU NUMBER = %d\n", runtime.GOARCH, runtime.GOOS, runtime.Version(), runtime.NumCPU() )
}


func array_prt() {
	fmt.Println("array_prt() begin .... ")

	var array_int []int = make([]int, 4)

	array_int[0] = 100
	array_int[1] = 200
	array_int[2] = 3
	array_int[3] = 400

	fmt.Println("数组循环输出 ... ")

	// 第一个返回值 ix 是数组或者切片的索引，第二个是在该索引位置的值；
	fmt.Println("两个返回值  ... ")
	for ix, value := range array_int {
		fmt.Printf("Slice[%d]: %d\n", ix, value)
	}

	// 忽略 index
	fmt.Println("一个返回值，忽略 index  ... ")
	for _, value := range array_int {
		fmt.Printf("Slice is: %d\n", value)
	}

	// 忽略 index
	fmt.Println("一个返回值，忽略 slice value  ... ")
	for ix := range array_int {
		fmt.Printf("Slice index : %d\n", ix)
	}

	defer fmt.Printf("array_prt() defer ... \n" )
	fmt.Println("array_prt() end .... ")
}

func array_cut() {
	fmt.Println("array_cut() begin .... ")

	arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("原数组：",arr)

	fmt.Println("对数组按 arr[1:4] 进行截取 ...")
	var b_array []int = arr[1:4] //creates a slice from a[1] to a[3]
 	fmt.Printf("截取后，数据是：%v; 长度：%d;  容量：%d\n" ,b_array, len(b_array), cap(b_array))
	fmt.Printf("数组截取之后的类型为：%T \n", b_array )

	//如果指定max，max的值最大不能超过截取对象（数组、切片）的容量
	//max:9  low：2  high;5  len:5-2(len=high-low)  cap:9-2(cap=max-low)
	s1:=arr[2:5:9]
	fmt.Printf("数据是：%v;长度：%d;  容量：%d\n" ,s1, len(s1), cap(s1))

	//如果没有指定max，max的值为截取对象（切片、数组）的容量
	s2:=s1[1:7]  //max:7  low：1  high;7  len:7-1(len=high-low)  cap:7-1(cap=max-low)
	fmt.Println("对切片进行截取：")
	fmt.Printf("对切片进行截取之后的数据是：%v,长度:%d； 容量%d\n",s2, len(s2), cap(s2))

}

func main() {
	fmt.Println("Hello, go slice!")
	two_array := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	fmt.Printf("二维数组 ：%v  \n " , two_array ) ;
	array_prt()
	array_cut()

}

