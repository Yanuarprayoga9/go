package main 

import (
	"fmt"
)

func main() {
	fmt.Println("if");

	a := 9;

	if a == 7 {
		fmt.Println("is 7")
	} else if a == 9 {
		fmt.Println("its 9")
	} else {
		fmt.Println("different")
	}
	
}