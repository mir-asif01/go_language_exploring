package main

import "fmt"

func main() {
	num := 10
	ptr := &num // &num refers to the address

	fmt.Println("pointer address is : ", ptr)
	fmt.Println("pointer value is ", *ptr) // *ptr refers to the value at ptr address
}
