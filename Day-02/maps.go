package main

import "fmt"

// Definisikan struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Inisialisasi slice dengan beberapa struct Person
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	// Menampilkan informasi dari setiap Person dalam slice
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}

	people2 := map[string]int{
		"test":1,
	}

	fmt.Println(people2["test"])
}
