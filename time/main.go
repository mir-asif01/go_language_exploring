package main

import (
	"fmt"
	"time"
)

func main() {
	_time := time.Now().Format("02-01-2006 Monday 15:04:05")
	// newTime := currentTime.Add(24*time.Hour)
	fmt.Println(_time)

	// timr_str := "31-12-2047"
	// parsedTime, _ := time.Parse("02-01-2006", timr_str)

	// fmt.Println("parsed time : ", parsedTime)
	// fmt.Printf("type of time_str is %t and type of parsedTime is %t", timr_str, parsedTime)
}
