package main

import "fmt"

func main() {
	// ! Map berfungsi untuk menyimpan satu atau lebih data namun, map disimpan sebagi key:value pairs
	// Map juga termasuk salah satu tipe data yang masuk dalam kategori reference type sama seperti tipe data slice.Berarti jika ada suatu map yang berusaha untuk meng-copy map lainnya, dan map tersebut mengganti value padasuatu key, maka value dari map lainnya tersebut juga akan ikut terganti.

	// Maksud dari inisialisasi di bawah berarti tipe data key dari map harus string dan value juga string.
	// * first way
	var person map[string]string //deklarasi
	person = map[string]string{} // insialisasi

	person["name"] = "Naufal"
	person["age"] = "21"
	person["address"] = "Jalan Gombel"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	// * second way
	var human = map[string]string{
		"name":    "Naufal",
		"age":     "21",
		"address": "Jalan Gombel",
	}

	fmt.Println("name:", human["name"])
	fmt.Println("age:", human["age"])
	fmt.Println("address", human["address"])

	// * Get map value with looping
	for key, value := range human {
		fmt.Println(key, ":", value)
	}

	// * Detect if value exist
	value, exist := human["age"]

	if exist {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	// * Delete map value
	delete(human, "age")

	fmt.Println("Human after deletion", human)

	// * Detect if value exist 2
	value2, exist2 := human["age"]

	if exist2 {
		fmt.Println(value2)
	} else {
		fmt.Println("Value has been deleted")
	}

	// * Combining map with slice
	var manusia = []map[string]string{
		{"name": "Naufal", "age": "21"},
		{"name": "Airell", "age": "23"},
		{"name": "Kiryu", "age": "60"},
	}

	for i, orang := range manusia {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, orang["name"], orang["age"])
	}
}
