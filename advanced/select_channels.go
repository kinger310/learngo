package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello, world"
}

func callerB(c chan string) {
	c <- "Hello, wangdi"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	for index := 0; index < 5; index++ {
		time.Sleep(1 * time.Microsecond)
		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("Default")
		}

	}
}
