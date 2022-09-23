package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ! Tipe data string pada Go terbentuk dari kumpulan tipe data byte yang di letakkan di dalam slice atau bisa kita sebutdengan slice of bytes.Tipe data byte pada Go merupakan tipe data alias dari tipe data uint8.

	word := "Cyberpunk"

	for i := 0; i < len(word); i++ {
		// Apabila stringnya dilooping kan mengembalikan bytesnya tuh semisal C = 67
		fmt.Printf("index: %d, bytes: %d \n", i, word[i])
	}

	// Setelah mengetahui hasil byte dari perulangan di atas, kita dapat membuat slice of byte untuk membentuk string tersebut kembali
	var cybunk []byte = []byte{67, 121, 98, 101, 114, 112, 117, 110, 107}
	fmt.Printf(string(cybunk))
	println()

	// * Ketika kita ingin mendapatkan panjang karakter dengan menggunakan function len(), maka sebetulnya kita tidak sedang mendapatkan panjang dari string berdasarkan karakternya, namun kita mendapatkan jumlah byte pada string.  Seperti contoh di bawah:
	str1 := "makan"
	str2 := "mânca"

	// Nah kan walau terlihat sama jumlah stringnya, tapi setelah dicek len-nya ternyata berbeda karena seperti deskripsi di atas.
	fmt.Printf("Str1 byte length => %d\n", len(str1))
	fmt.Printf("Str2 byte length => %d\n", len(str2))

	// ! Jikalau ingin mencari jumlah karakternya, bukan bytenya maka perlu merubah string tersebut menjadi rune terlebih dahulu seperti di bawah baru benar:
	fmt.Printf("Str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("Str2 character length => %d\n", utf8.RuneCountInString(str2))

	// ! atau bisa juga menggunakan range loop untuk loop range-per-range, dan bahkan pada loop ini index pada looping langsung lompat 2 dari 1 ke 3, ini karena â terdiri dari 2 byte.
	for index, value := range str2 {
		fmt.Printf("index => %d, string => %s\n", index, string(value))
	}

}
