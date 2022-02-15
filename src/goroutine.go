package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1)
	runtime.Gosched()
	i := 0
	go func() {
		for ; i < 100; i++ {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)

		}
	}()

	time.Sleep(time.Millisecond * 1000)
	i = 10

}
