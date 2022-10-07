package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Employee struct {
	FullName string `json:"full_name"` // tulisan2 json: disamping merupakan tag yang digunakan untuk decode data2 seperti json, form, xml/
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	// * Decode JSON to struct
	// Kita akan mendecode sebuah data json di bawah kepada sebuah struct.
	var jsonString = `
		{
			"full_name": "Naufal Aldy Pradna",
			"email": "naufal@gmail.com",
			"age": 21
		}		
	`

	var result Employee

	// Codeline dibawah mendecode JSON menjadi struct employee, dengan argumen pertama berupa slice of byte dari json, dan argumen kedua berupa pointer dari variable result.
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("JSON to Struct")
	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)

	// * Decode json to map yang memiliki key berupa string, dan interface kosong untuk mendapatkan nilai dengan tipe data berbeda2
	var result2 map[string]interface{}

	var err2 = json.Unmarshal([]byte(jsonString), &result2)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("JSON to Map")
	fmt.Println(strings.Repeat("#", 50))
	fmt.Println("full_name:", result2["full_name"])
	fmt.Println("email:", result2["email"])
	fmt.Println("age:", result2["age"])

	// * Decode json to empty interface
	var temp interface{}

	var err3 = json.Unmarshal([]byte(jsonString), &temp)
	if err3 != nil {
		fmt.Println(err.Error())
		return
	}
	// Untuk mengakses data didalam harus di-type assertation menjadi map dengan key string dan value empty interface
	var result3 = temp.(map[string]interface{})

	fmt.Println("JSON to Empty Interface")
	fmt.Println(strings.Repeat("#", 50))
	fmt.Println("full_name:", result3["full_name"])
	fmt.Println("email:", result3["email"])
	fmt.Println("age:", result3["age"])

	// * Decode json array to slice of struct
	var jsonArray = `[
		{
			"full_name": "Naufal Aldy Pradna",
			"email": "naufal@gmail.com",
			"age": 21
		},
		{
			"full_name": "Airella",
			"email": "airell@gmail.com",
			"age": 23
		}				
	]`

	var result4 []Employee

	var err4 = json.Unmarshal([]byte(jsonArray), &result4)
	if err4 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("JSON Array to Slice of Struct")
	fmt.Println(strings.Repeat("#", 50))
	for i, v := range result4 {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}
