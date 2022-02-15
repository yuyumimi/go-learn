package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程 i=", i, len(ch), cap(ch))
			ch <- i
		}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主go成 num=", num)
	}
}
