package main

import (
	"fmt"
	"time"
)

type Book struct {
	title   string
	author  string
	subject string
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println(111)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("word")
	say("hello")
}

func swap(x, y int) {
	fmt.Println(x)
	fmt.Println(y)
}
