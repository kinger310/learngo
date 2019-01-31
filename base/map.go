package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bells"] = Vertex{111, 222}
	m["Facebook"] = Vertex{232, 4324.0}
	fmt.Println(m, m["Bells"], m["bells"])
}

// func printSlice(s []) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }
