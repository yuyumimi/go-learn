package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func print(s string) {
	mutex.Lock()
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(100 * time.Millisecond)
	}
	mutex.Unlock()
}

func main() {
	go person1()
	go person2()
	go person3()

	time.Sleep(time.Second * 3)
}

func person1() {
	print("hello")
}
func person2() {
	print("world")
}
func person3() {
	print("zhangsan")
}
