package main

import "fmt"


func main() {
	nums := []int{1, 2, 3, 4, 5}

	for index, values := range nums {
		fmt.Println("index",index)
		fmt.Println("value",values)
	}
}