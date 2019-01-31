package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	printSlice(slice)
	newSlice := append(slice, 5)
	printSlice(newSlice)

	// 使用...运算符，可以将一个切片的所有元素追加到另一个切片里
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5}
	s := append(s1, s2...)
	printSlice(s1)
	printSlice(s2)
	printSlice(s)

	// 	For slice[i:j:k] or [2:3:4]
	// Length: j - i or 3 - 2 = 1
	// Capacity: k - i or 4 - 2 = 2
	source := []int{1, 2, 3, 4, 5}
	printSlice(source)
	s3 := source[2:3:3]
	s4 := source[2:3]
	printSlice(s3)

	// 设置切片长度和容量一致的好处
	// if s3 := source[2:3:4] source会发生改变；else if 3 := source[2:3:3], source 不会变。
	// 不要动底层数组
	s3 = append(s3, 6)
	printSlice(source)
	s4 = append(s4, 7)
	printSlice(source)
	printSlice(s3)
	printSlice(s4)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
