package main

import "fmt"

func main() {
	fmt.Println("declarations array")
	var pointerArray [3]*int
	num := 1
	num1 := 2
	num2 := 3

	pointerArray[0] = &num
	pointerArray[1] = &num1
	pointerArray[2] = &num2
	for i := 0; i < len(pointerArray); i++ {
		fmt.Println(*pointerArray[i])
	}

	type Person struct {
		name string
		age  uint
	}

	persons := []Person{
		{
			name: "yanuar",
			age:  6,
		}, {
			name: "dimas",
			age:  6,
		},
	}

	for i := 0; i < len(persons); i++ {
		fmt.Println("nama = ", persons[i].name)
		fmt.Println("age = ", persons[i].age)
	}

	for _, person := range persons {
		fmt.Println("nama = ", person.name)
		fmt.Println("age = ", person.age)
	}

	searchName := func(name string, persons []Person) []Person {
		var result []Person

		for _, person := range persons {
			if person.name == name {
				result = append(result, person)
			}
		}
		return result

	}

	yanuar := searchName("yanuar", persons)

	fmt.Println(yanuar)
}
