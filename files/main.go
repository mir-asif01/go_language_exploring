package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Create("test.txt")
		if err != nil {
			fmt.Println("Error : ", err)
			return
		}
		fmt.Println("File created successfuly")

		code := "<h1>Hello go, I am creating by you</h2>"
		_, error := io.WriteString(file, code)
		if error != nil {
			fmt.Println("error while writting into file", error)
		}
	*/

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println("error while reading", err)
	// }
	// defer file.Close()

	// buffer := make([]byte, 1024)

	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		fmt.Println("End of file")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Eorro in buffer", err)
	// 		return
	// 	}

	// 	fmt.Println(string(buffer[:n]))
	// }

	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error while reading file", err)
	}
	fmt.Println(string(file))

}
