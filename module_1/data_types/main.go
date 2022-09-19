package main

import "fmt"

func main() {
	// Ada baiknya saat deklarasi variabel menggunakan jenis spesifik yang diinginkan untuk menghindari boros memori

	// Number (Integer)
	var first uint8 = 98  // uint memiliki range 0 - 255
	var second int8 = -12 // int memiliki range -128 - 127

	// Cek tipe data
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", second)

	// Number (Float) ada float32 dan float64
	var decimalNum float32 = 3.14

	// Pada dasarnya %f memformat angka desimal menjadi 10^-6, sehingga diberikan .3 untuk menyesuaikan output
	fmt.Printf("decimal number: %f\n", decimalNum)
	fmt.Printf("decimal number: %.3f\n", decimalNum)

	// Bool cuma true false as usual
	var apakahBenul = false
	fmt.Println("Apakah benul?", apakahBenul)

	// String
	// Coba string panjang
	var message = `When i see you smile i
	can face za warudo. 
	owoooooo`

	fmt.Println(message)

	/*
		nil dan zero value, pada golang semua zero value(deklarasi tapi tidak diberi nilai) tetap ada default valuenya
			● Zero value dari string adalah "" (string kosong).
			● Zero value dari bool adalah false.
			● Zero value dari tipe numerik non-desimal adalah 0.
			● Zero value dari tipe numerik desimal adalah 0.0.

			Zero value berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong. Nil tidak bisa digunakan pada tipe datayang sudah dibahas sebelumnya. Ada beberapa tipe data yang bisa di-set nilainya dengan nil, diantaranya:
			●pointer
			●tipe data function
			●slice●map
			●channel
			●interface kosong atau interface{}
	*/

}
