package main 

import (
	"fmt"
)

func main () {
	fmt.Println("type aliases");
	type test int64

	var x test = 1024

	fmt.Println(x);
}