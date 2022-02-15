package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		send(ch)
	}()

	rec(ch)
}

func send(out chan<- int) {
	out <- 888

}

func rec(in <-chan int) {
	num := <-in
	fmt.Println(num)
}
