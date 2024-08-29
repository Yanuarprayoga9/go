package main

import "fmt"

func main() {
	fmt.Println("string:")
	n  := make(chan int)
	// n := "Hello, World!" // this also work as short hand declaration
	fmt.Println("n=", n)
}