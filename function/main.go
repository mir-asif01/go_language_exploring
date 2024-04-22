package main

import "fmt"

func addNumber(n1, n2 int) int {
	return n1 + n2
}

func variadicFunc(values ...int) {
	for _, v := range values {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("Learning functions in golang")
	// fmt.Println("Sum :: ", addNumber(2, 3))
	variadicFunc(1, 23, 4, 5, 6)
}
