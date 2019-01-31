package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 分配一个
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				time.Sleep(50 * time.Millisecond)
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")

	wg.Wait()
	fmt.Println("\nTerminating Program")
}
