package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Is_Adult bool   `json:"is_adult"`
}

func main() {

	person := Person{Name: "Asif", Age: 21, Is_Adult: true}
	fmt.Println("Person befor marshalling", person)

	//json encoding {marshalling}
	mrslPerson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error is:", err)
	}
	fmt.Println("Marshaled person :", string(mrslPerson))

	//json decoding {unmarshaling}
	var unmrslPerson Person
	err = json.Unmarshal(mrslPerson, &unmrslPerson)
	if err != nil {
		fmt.Println("Error on unmarshaling", err)
	}
	fmt.Println("Unmarshaled", unmrslPerson)

}
