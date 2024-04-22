package main

import "fmt"

func main() {
	fmt.Println("Learning array in golang!")
	nums := [5]int{1, 3, 5, 7, 9} // defining an array
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	for index, value := range nums {
		fmt.Printf("Value is %d and index is %d\n", value, index)
	}
}
