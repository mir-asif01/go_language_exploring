package main

import "fmt"

// main funtion
func main() {
	add := closureFunc() // initializing closure function
	fmt.Println(add(5))
	fmt.Println(add(20))
	fmt.Println(add(5))

}

// this closure function defination
func closureFunc() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
