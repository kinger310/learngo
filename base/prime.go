package main

import "fmt"

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n%3 == 0 {
		return false
	}
	i := 5
	w := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i += w
		w = 6 - w
	}
	return true
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i, isPrime(i))
	}
}
