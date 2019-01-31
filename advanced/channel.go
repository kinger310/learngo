package main

import (
	"fmt"
	"time"
)

// The example code sums the numbers in a slice, distributing the work between two goroutines.
// Once both goroutines have completed their computation, it calculates the final result.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(1 * time.Second)
	c <- sum // send sum to channel c
}

func main() {
	for index := 0; index < 20; index++ {
		s := []int{1, 2, 3, 4, 5, 6, 7}
		c := make(chan int, 2)

		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
		// x, y := <-c, <-c //receive from channel c
		// fmt.Println(x, y, x+y)
		fmt.Println(<-c, <-c)
	}

}
