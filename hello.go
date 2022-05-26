package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(add(1, 2))
}

func add(a, b int) (int) {
	sum := a + b
	return sum;
}