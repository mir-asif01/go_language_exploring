package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello golang")
	// var day string = string(time.Now().Weekday())
	// day := "Friday"

	// fmt.Println(day)

	// switch day {
	// case "Sunday":
	// 	fmt.Println("Weekend")
	// case "Friday":
	// 	fmt.Println("Weekend")
	// 	fallthrough
	// default:
	// 	fmt.Println("Workday")
	// }

	/// clojures///#####################

	add := myFunc()
	add(5)
	fmt.Println(add(5))

}

func myFunc() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
