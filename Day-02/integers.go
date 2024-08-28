package main

import "fmt"

func main () {
	fmt.Println("signed integer");
	var a int8 = 127 //range -128 to 127
	var b int16 = -32768 //range -32768 to 32767
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("unsigned integer");
	var unsignedIntA uint8 = 255 // ranges 0 to 255
	var unsignedIntB uint16 = 255 // ranges 0 to 65k
	fmt.Println("unsignedIntA = ",unsignedIntA)
	fmt.Println("unsignedIntB = ",unsignedIntB)


	fmt.Println("Machine depedent types: ")
	var mce int = 123456789
	var mcf uint = 123456789
	var mcg uintptr = 0xdeadbeef
	fmt.Println("mce = ",mce)
	fmt.Println("mcf = ",mcf)
	fmt.Println("mcg = ",mcg)
}