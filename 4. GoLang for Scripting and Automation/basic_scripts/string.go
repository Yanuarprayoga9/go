package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "halo"
	str2 := " yanuar"
	fmt.Println(str+str2)

	sentence := "Go is an open source programming language";
	words := strings.Split(sentence, " ")
	fmt.Println("Words in sentence:", words)
}