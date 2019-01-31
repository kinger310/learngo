package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// init在main函数之前调用
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage go run <file> <url>")
		os.Exit(-1)
	}
}

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return

	}
	io.Copy(os.Stdout, r.Body)
	if err := r.Body; err != nil {
		fmt.Println(err)
	}
}
