package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://developer.com:80/hello?name=airell&age=23"
	// * Melakukan parsing terhadap url string menjadi url.Url, dimana kita kemudian bisa mengakses banyak hal seperti protocol, host, query, dkk.
	var u, err = url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s \n", name, age)
}
