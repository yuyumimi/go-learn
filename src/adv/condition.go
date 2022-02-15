package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	//生成随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 15; i++ {
		go readGo(ch, i)
	}
	for i := 0; i < 15; i++ {
		go writeGo(ch, i)
	}
	time.Sleep(time.Second * 3)
}

func writeGo(ch chan<- int, i int) {
	value := rand.Intn(1000)
	fmt.Printf("%dth,write: %d\n", i, value)
	ch <- value
}

func readGo(ch <-chan int, i int) {
	num := <-ch

	fmt.Printf("%dth,read: %d\n", i, num)

}
