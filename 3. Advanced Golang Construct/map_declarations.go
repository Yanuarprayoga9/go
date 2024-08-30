package main

import "fmt"

type Person struct {
	name string
	age uint
}

type Responses struct {
	total uint
	persons []Person
}


func main() {
	var mapSlice []map[string]int

	map1 := map[string]int{"a": 1, "b": 2}

	mapSlice = append(mapSlice, map1)

	fmt.Println(mapSlice[0]["a"]);


	fruits := map[string]int{"apple":1,"banana":7};
	fruits["orange"] = 10;
	// Checking Existence
	price, exists := fruits["orange"]
	if exists {
		fmt.Println("Price of kiwi:", price)
	} else {
		fmt.Println("Kiwi does not exist in the map")
	}
	
}