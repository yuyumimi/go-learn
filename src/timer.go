package main

import (
	"fmt"
	"time"
)

func main() {

	//fmt.Println("now: ", time.Now())

	//timer := time.NewTimer(time.Millisecond * 3000)
	//ch := <-timer.C
	//fmt.Println(ch)

	//after := time.After(time.Second * 3)
	//
	//time.Sleep(time.Second * 4)
	//fmt.Println("now:  ", time.Now())
	//
	//ch := <-after
	//fmt.Println(ch)

	i := 0
	mytimer(5, func() {
		i++
		fmt.Println("执行任务。。。。", i)
	})
}

func mytimer(t int, f func()) {

	for {

		fmt.Println("begin:  ", time.Now())

		after := time.After(time.Second * time.Duration(t))

		f()
		times := <-after
		fmt.Println("end:  ", time.Now())
		fmt.Println(times)

	}
}
