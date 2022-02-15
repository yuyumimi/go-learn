package main

import "fmt"

func main() {
	var ch = make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程 i=", i)
			ch <- i
		}
	}()

	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主go成 num=", num)
	}
}
