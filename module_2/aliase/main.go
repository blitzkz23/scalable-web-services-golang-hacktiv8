package main

import "fmt"

func main() {
	// ! Aliase merupakan sebuah fitu pada bahasa Go yang digunakan sebagai nama alternative dari tipe data yangsudah ada.

	// Tipe data byte merupakan tipe data aliase dari tipe data uint8, yang berarti mereka merupakan tipe data yangsama dengan nama yang berbda.
	// Deklarasi variable dengan tipe data uint8
	var a uint8 = 10
	var b byte // byte adalah alias dari tipe data uint8

	b = a // tidak error karena memang aliasnya
	_ = b

	fmt.Println("=============================")
	// Delkarasi tipe data (baru) alias bernama second
	// type nama_alias = nama_tipe_data
	type second = uint

	// Variable hour bertipe data second dimana aslinya merupakan sebuah uint
	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint
}
