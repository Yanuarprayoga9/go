package main

import "fmt"


type Person struct {
	Name string
	Age uint16
}

func changeName (person *Person) {
	person.Name = "editr";
}
func main() {
	fmt.Println("struct:")
	person1 := Person{"yanuar",21}
	var person1ptr  *Person = &person1 ;
	fmt.Printf("%p",person1ptr)
	changeName(&person1);
	fmt.Println(person1ptr)}