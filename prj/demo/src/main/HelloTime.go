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

	endTime := time.Since(t0) // 耗时
	fmt.Println("init() 运行时间 : ", endTime)

}

func main() {
	fmt.Println("Hello, go TIME ")

	fmt.Println("时间戳 10位秒，13位毫秒 13位纳秒 ")
	// 10位数的时间戳是以 秒 为单位；
	// 13位数的时间戳是以 毫秒 为单位；
	// 19位数的时间戳是以 纳秒 为单位；
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n",time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n",time.Now().UnixNano() / 1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n",time.Now().UnixNano() / 1e9)

	fmt.Println("格式化输出 ... ")
	t := time.Now()
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%04d-%02d-%02d\n", t.Year() , t.Month(), t.Day() )
	// 输出纳秒
	fmt.Printf("输出纳秒 ：%04d-%02d-%02d %02d:%02d:%02d %03d\n", t.Year() , t.Month(), t.Day(),t.Hour(),t.Minute(),t.Second(),t.Nanosecond())
	// 纳秒转成毫秒
	fmt.Printf("输出毫秒 ：%04d-%02d-%02d %02d:%02d:%02d %03d\n", t.Year() , t.Month(), t.Day(),t.Hour(),t.Minute(),t.Second(),t.Nanosecond()/1000000)
	// 21.12.2011
	//t = time.Now().UTC()
	//fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
	//fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	//var  week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	//week_from_now := t.Add(time.Duration(week))
	//fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	//fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	//fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011

	println("layout ... ");
	//fmt.Println(t.Format("02 Jan 2006 15:04:02")) // 21 Dec 2011 08:52
	//s := t.Format("2006-01-02")
	//fmt.Println("--- ", t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221

	// 时间 to 时间戳
	//设置时区
	println("时间 --> 时间戳 ... ");
	loc, _ := time.LoadLocation("Asia/Shanghai")
	//2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-11 15:07:51", loc)
	fmt.Println("Unix time : " ,tt.Unix())

	// 时间戳 to 时间
	println("时间戳 --> 时间 ... ");
	tm := time.Unix(1531293019, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19

	// tm = time.Now() ;
	// fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19
	// fmt.Println(t.Unix().Format("2006-01-02 15:04:05"))
}

