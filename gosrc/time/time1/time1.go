// time1
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	t0 := time.Date(2015, 11, 18, 7, 56, 35, 0, time.UTC)
	fmt.Println(t0) // 2011-11-18 07:56:35 +0000 UTC

	t := time.Now()
	fmt.Println(t)                        // 2016-08-04 11:42:16.2741028 +0800 CST
	fmt.Println("Year =", t.Year())       // 2016
	fmt.Println("Month =", t.Month())     // August
	fmt.Println("YearDay =", t.YearDay()) // 217
	fmt.Println("Day =", t.Day())         // 4
	fmt.Println(time.Now().Date())        // 2016 August 4
	fmt.Println("---------------------------------")

	fmt.Println("日期格式化")
	// 格式化：format.go
	fmt.Println(t.Format(time.ANSIC)) // Thu Aug  4 11:51:48 2016
	layout := "01__02-2006 3.04.05 PM"
	fmt.Println(time.Now().Format(layout))                       // 08__04-2016 11.51.48 AM
	fmt.Println(time.Now().Format("20060102150405000"))          // 20160804134249000
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))    // 2016-08-04 13:42:18.409
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000")) // 2016-08-04 13:43:34.120325
	fmt.Println("---------------------------------")

	fmt.Println("日期计算")
	// Add方法和Sub方法是相反的，获取t0和t1的时间距离d是使用Sub
	tNow := time.Now()
	// 一天之前
	d, _ := time.ParseDuration("-24h")
	fmt.Println("d type:", reflect.TypeOf(d))
	fmt.Println("d =", d)
	fmt.Println("一天之前:", tNow.Add(d)) // 一天之前: 2016-08-03 14:23:03.5289553 +0800 CST
	// 一周之前
	fmt.Println("一周之前:", tNow.Add(d*7)) // 一周之前: 2016-07-28 14:23:03.5289553 +0800 CST
	// 一月之前
	fmt.Println("一月之前:", tNow.Add(d*30)) // 一月之前: 2016-07-05 14:23:03.5289553 +0800 CST
	preMonth := time.Date(tNow.Year(), tNow.Month()-1, tNow.Day(), tNow.Hour(), tNow.Minute(), tNow.Second(), tNow.Nanosecond(), tNow.Location())
	fmt.Println("一月之前:", preMonth) // 一月之前: 2016-07-04 14:23:03.5289553 +0800 CST

	fmt.Println(tNow.Sub(t0))                 // 6238h18m46.719706s
	fmt.Println(reflect.TypeOf(tNow.Sub(t0))) // time.Duration
	fmt.Println("---------------------------------")

	fmt.Println("Before & After方法") // 判断一个时间点是否在另一个时间点的前面（后面），返回true或false
	t1 := time.Now()
	time.Sleep(time.Second)
	t2 := time.Now()
	a := t2.After(t1)  //t2的记录时间是否在t1记录时间的**后面**呢，是的话，a就是true
	fmt.Println(a)     //true
	b := t2.Before(t1) //t2的记录时间是否在t1记录时间的**前面**呢，是的话，b就是true
	fmt.Println(b)     //false
	fmt.Println("---------------------------------")

	// 以下参考：
	// go语言的time包：http://my.oschina.net/u/943306/blog/149395
	fmt.Println("Sleep函数") // 表示睡多少时间，睡觉时，是阻塞状态
	fmt.Println("start sleeping...")
	time.Sleep(time.Second)
	fmt.Println("end sleep.")
	fmt.Println("---------------------------------")

	fmt.Println("After函数") // 和Sleep差不多，意思是多少时间之后，但在取出管道内容前不阻塞
	fmt.Println("the 1")
	tc := time.After(time.Second) //返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点(time.Now())
	//时间点记录的是放入管道那一刻的时间值
	fmt.Println("the 2")
	fmt.Println("the 3")
	<-tc //阻塞中，直到取出tc管道里的数据
	fmt.Println("the 4")
	fmt.Println("---------------------------------")

	fmt.Println("AfterFunc函数") // 和After差不多，意思是多少时间之后在goroutine line执行函数
	fmt.Println("start ...")
	f := func() {
		fmt.Println("Time out")
	}
	time.AfterFunc(1*time.Second, f)
	fmt.Println("after AfterFunc ...")
	time.Sleep(2 * time.Second) //要保证主线比子线“死的晚”，否则主线死了，子线也等于死了
	fmt.Println("---------------------------------")

	fmt.Println("Tick函数") // 和After差不多，意思是每隔多少时间后，其他与After一致
	fmt.Println("the 1")
	tc = time.Tick(time.Second) //返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点，
	//1秒后再放一个，一直反复，时间点记录的是放入管道那一刻的时间
	for i := 1; i <= 2; i++ {
		<-tc
		fmt.Println("hello")
	} //每隔1秒，打印一个hello
}
