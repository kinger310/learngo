package main

import (
	"fmt"
	"time"
)

func prod(c chan int) {
	for index := 0; index < 5; index++ {
		// fmt.Println("Produce >>", index)
		c <- index
		fmt.Println("Produce >>", index)
	}
}

func consume(c chan int) {
	for index := 0; index < 5; index++ {
		num := <-c
		fmt.Println("Consume <<", num)

	}
}

func main() {
	c := make(chan int, 5)
	go consume(c)
	go prod(c)
	time.Sleep(100 * time.Millisecond)

}
