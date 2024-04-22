package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello golang")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Printf("Your full name is %s", name)
}
