package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwmutex sync.RWMutex
var value int

func main() {
	fmt.Printf("@@@@@@@%d", value)
	//生成随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 15; i++ {
		go readGo(i)
	}
	for i := 0; i < 5; i++ {
		go writeGo(i)
	}
	time.Sleep(time.Second * 3)
}

func writeGo(i int) {
	rwmutex.Lock()
	value = rand.Intn(1000)
	fmt.Printf("%dth,write: %d\n", i, value)
	rwmutex.Unlock()
}

func readGo(i int) {
	rwmutex.RLock()

	fmt.Printf("%dth,read: %d\n", i, value)
	rwmutex.RUnlock()

}
