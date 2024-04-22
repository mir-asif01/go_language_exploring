package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("learning maps in go!!")
	marks := make(map[string]int)

	marks["rom"] = 10
	marks["seth"] = 10
	marks["john"] = 30

	// check if the key is present in the map or not
	// _, prs := marks["seth"]

	// fmt.Println("prs:", prs)
	// clears the map
	// clear(marks)

	// delete single item from map
	delete(marks, "john")
	// fmt.Println("Mark :", marks)

	n1 := map[string]string{"name": "asif", "age": "21"}
	// fmt.Println(n1)

	n2 := map[string]string{"name": "asif", "age": "22"}
	// fmt.Println(n1)

	isEqul := maps.Equal(n1, n2)
	fmt.Println("value of equal is :", isEqul)

}
