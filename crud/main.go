package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getMethod() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	if err != nil {
		fmt.Println("Error while get request : ", err)
		return
	}
	post, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading ioutil", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in response :", res.StatusCode)
	}

	fmt.Println("Post body is : ", string(post))
	defer res.Body.Close()
}

func postMethod() {
	newPost := Post{UserId: 199, Id: 19, Title: "New post title", Body: "new post body post request by me"}

	marshaledPost, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println(err)
		return
	}
	postString := string(marshaledPost)
	postReader := strings.NewReader(postString)

	url := "https://jsonplaceholder.typicode.com/posts"

	res, err := http.Post(url, "application/json", postReader)

	if err != nil {
		fmt.Println("error", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return
	}
	fmt.Println(res.StatusCode)
}

func updateMethod() {
	newPost := Post{UserId: 199, Id: 190, Title: "New post title", Body: "new post body post request by me"}

	marshaledPost, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println(err)
		return
	}
	postString := string(marshaledPost)
	postReader := strings.NewReader(postString)

	url := "https://jsonplaceholder.typicode.com/posts/2"

	req, err := http.NewRequest("PUT", url, postReader)
	req.Header.Set("content-type", "application/json")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return
	}
	fmt.Println(res.Status)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Res body :", string(data))
	defer res.Body.Close()
}

func deleteMethod() {
	url := "https://jsonplaceholder.typicode.com/posts/2"
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error while req", err)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while res", err)
	}
	deletedData, err := ioutil.ReadAll(res.Body)

	fmt.Println("status of response", res.Status)
	fmt.Println("body of response", string(deletedData))

	defer res.Body.Close()
}

func main() {
	// postMethod()
	getMethode()
	// updateMethod()
	fmt.PrintLn("Learning in Lapce code editor !!!")
	// deleteMethod()
}
