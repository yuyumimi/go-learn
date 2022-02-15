package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	ticker := time.NewTicker(time.Second * 1)

	go func() {
		i := 0
		for {
			c := <-ticker.C
			fmt.Println(c)

			i++
			if i == 3 {
				ch <- 111
				break
			}

		}
	}()

	<-ch
}
