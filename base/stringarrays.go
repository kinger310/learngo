package main

import "fmt"

func main() {
	// string arrays
	names := [4]string{
		"andy",
		"bebe",
		"cindy",
		"dony",
	}
	fmt.Println(names)
	// slice
	a := names[:3]
	fmt.Println(a)
	// arrays append elements
	// https://tour.golang.org/moretypes/15
	// names_new := "ella"
	// names_new := []string{
	// 	"ella",
	// }
	// names = append(names, names_new)

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}

	var ss []string
	printSlice(ss)
	ss = append(ss, "0")
	printSlice(ss)
	ss = append(ss, "1", "2", "3")
	printSlice(ss)

	// slice := []int{10, 20, 30, 40, 50}
	// newSlice := slice[1:3]
	// printSlice(newSlice)
	// newSlice[1] = 1
	// printSlice(newSlice)
}

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
