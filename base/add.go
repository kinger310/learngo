package main

import "fmt"

func add(count int) int {
	sum := 0
	for index := 0; index < count; index++ {
		sum += index
		fmt.Println(sum)
	}
	return sum
}

func main() {
	fmt.Println(add(10))
}
