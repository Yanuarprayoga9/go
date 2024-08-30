package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	yanuar := Person{Name: "Yanuar", Age: 21}

	jsonData, err := json.Marshal(yanuar)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonData)

	var unMarshalData Person
	 err = json.Unmarshal(jsonData, &unMarshalData)
	 if err != nil  {
		log.Fatal(err)
	 }
	 fmt.Println(unMarshalData)
}