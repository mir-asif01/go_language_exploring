package main

import (
	"fmt"
	"time"
)

func hiFunc() {
	fmt.Println("from hiFunc.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("from hiFunc ended.")
}

func helloFunc() {
	fmt.Println("from helloFunc.")
}

func main() {
	fmt.Println("Learning go routines......")
	go hiFunc()
	helloFunc()
	time.Sleep(1000 * time.Millisecond)
}
