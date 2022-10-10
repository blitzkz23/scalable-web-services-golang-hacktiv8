package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	// ! Coba GET Request
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ini response body berupa pointer dari struct http Response:", response.Body)
	// ReadAll() akan mengembalikan slice of byte dan error.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Di defer agar terakhri dijalankan.
	defer response.Body.Close()

	sb := string(body)
	fmt.Println("Ini adalah isi body berupa slice of byte yang sudah di convert ke string:", sb)

	fmt.Println(strings.Repeat("=", 50))
	// ! Coba POST Request
	data := map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}

	var stringPayload = `
		[
			{
				"title":  "foo",
				"body":   "bar",
				"userId": 1,
			},
			{
				"title":  "foo",
				"body":   "bar",
				"userId": 1,
			}
		]
	`
	_ = stringPayload

	// Marshal() akan mengembalikan slice of byte dan error.
	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatal(err)
	}

	// * Membuat request ke API menggunakan data yang telah diubah tipe datanya menjadi interface io.Reader karena parameter ketiga dari http.NewRequest() adalah interface io.Reader berupa body request.
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	// * client.Do untuk eksekusi request.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body2, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body2))
}
