package main

import "fmt"

type Hasname interface {
	GetName() string
	SetName()
}

type Animal struct {
	name string
}

func (anim *Animal) GetName() string {
	return anim.name
}
// func (anim *Animal) SetName(s string) {
// 	anim.name = s
// }
func main() {
	var cat Animal;	
	// cat.SetName("cat")
	fmt.Println(cat.GetName())

}