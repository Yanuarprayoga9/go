package main

import "fmt"

func main() {
	fmt.Println("pointer:")
	var n string = "Hello, World!"
	var nptr *string = &n;

	// n := "Hello, World!" // this also work as short hand declaration
	fmt.Println("n=", nptr,n,*nptr)
}