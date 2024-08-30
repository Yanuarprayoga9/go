package main

import "fmt"

func endApp () {
	fmt.Println("APLIKASI SELESAI")
	message := recover();
	fmt.Println(message)
}

func runApp(error bool) {
	defer endApp();
	if error{
		panic("ERROR")
	}
}
func main() {
	runApp(true);
}