package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	timeout()
	return

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		ch <- 1
		ch <- 3
		ch <- 2
		close(ch)
		quit <- true
		runtime.Goexit()
	}()

	for {
		select {
		case num := <-ch:
			fmt.Println("读到数据：", num)
		case <-quit:
			return
		}

		fmt.Println("====")
	}

}

func timeout() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		time.Sleep(time.Second * 5)
		for {
			select {
			case num := <-ch:
				fmt.Println("读到数据：", num)
			case <-time.After(3 * time.Second):
				//fmt.Println("after...")
				quit <- true
				goto label
			}

		}
	label:
		fmt.Println("label====")
	}()

	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

	<-quit
	fmt.Println("over!!!")
}

func fibonacci(n int) int {

	x, y := 1, 1
	for i := 0; i < n; i++ {
		fmt.Print(x, " ")
		x, y = y, x+y
	}

	return x
}
