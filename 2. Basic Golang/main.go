package main

import (
	"day2/helper"
	"day2/shapes"
	"fmt"

	"github.com/google/uuid"
)

func main () {
	shapes.AreaOfCircle(3)
	helper.PrintHello();

	test := uuid.New().String()

	fmt.Println(test)

	
}