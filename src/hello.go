package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	//stdout := os.Stdout
	stderr := os.Stderr
	str := "hello"
	fmt.Println("who is a" == "who")
	stderr.Write(bytes.NewBufferString(str).Bytes())

}
func test1() *Person {

	p := Person{"yu", 1, 18}
	return &p
}
func test(man *Student) {
	(*man).person.age = 12
	man.course = "chinese"
	fmt.Println(man)
}

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	person Person
	course string
}
