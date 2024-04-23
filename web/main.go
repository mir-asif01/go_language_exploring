package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning about go web requests")
	res, _ := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	post, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(post))

	defer res.Body.Close()
}
