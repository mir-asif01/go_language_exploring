package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&t=6s"

	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error while parsing", err)
	}

	fmt.Println("parsed url", parsedUrl)
	// fmt.Printf("type of parsed url is %T", parsedUrl)
	// manuy parts of url
	/*
		fmt.Println("Schema is : ", parsedUrl.Scheme)
		fmt.Println("host is : ", parsedUrl.Host)
		fmt.Println("path is :", parsedUrl.Path)
		fmt.Println("query", parsedUrl.RawQuery)
	*/

	parsedUrl.Path = "myPathUrl"
	parsedUrl.RawQuery = "email=asif@gmail.com"

	newUrl := parsedUrl.String()

	fmt.Println("new url is :", newUrl)

}
