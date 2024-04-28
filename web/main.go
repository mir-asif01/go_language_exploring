package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type to_do struct {
	userId string `json:userId`
	id     string `json:id`
	title  string `json:title`
	body   string `json:body`
}

func main() {
	fmt.Println("Learning about go web requests")
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	if err != nil {
		fmt.Println("Error", err)
	}
	var newTodo to_do
	post, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(post))

	defer res.Body.Close()
}
