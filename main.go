package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(Sum(1, 2))
}

func Sum(a, b int) int {
	return a + b
}
