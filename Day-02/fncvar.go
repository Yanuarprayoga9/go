package main 

import (
	"fmt"
)



func main () {
	fmt.Println("function as variale")

	test := func (id int) (string,int) {
		if id == 1 {
			return "test",1
		}
		return "halo",1
	}
	fmt.Println(test(2));


}