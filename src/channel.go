package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func print(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go person1()
	go person2()
	go person3()

	<-ch
}

func person1() {
	print("hello")
	ch <- 88
}
func person2() {
	<-ch
	print("world")
}
func person3() {
	<-ch
	print("zhangsan")
}
