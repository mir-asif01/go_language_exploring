package main

import "fmt"

func main() {
	// var s []string
	// s = make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	s := []string{"a", "b", "c", "d"}

	fmt.Println(s)

	fmt.Println("After appending")
	s = append(s, "e", "f")
	fmt.Println(s)
}
