package main

import "fmt"

func main() {
	fmt.Println("array: static size")
	var o [6]int = [6]int{1,2,3,4,5,6};
	for a:=0 ; a < len(o) ; a++ {
		fmt.Println("n=", o[a])
	}
	// n := "Hello, World!" // this also work as short hand declaration
	fmt.Println("n=", 0)
}