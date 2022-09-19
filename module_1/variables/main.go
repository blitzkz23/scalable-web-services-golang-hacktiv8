package main

import "fmt"

// Cara inisialisasi seperti ini bagus untuk dijadikan variable global agar var/const tidak redundant
var (
	a int8   = 20
	b string = "hello"
)

var sebuahString = fmt.Sprintf("Ini gunanya untuk menampung nilai print %s", b)

func main() {
	_, _, _ = a, b, sebuahString
	// * Variable with data types
	var name string = "Naufal"
	var age int = 21

	fmt.Println("Nama gweh", name, "dengan umur", age)

	// * Variable with late assignment
	var name2 string
	name2 = "Aldy"

	fmt.Println("My another name", name2)
	println()

	// * Variable without data type (type inference)
	var apakahBenul = true

	fmt.Println("Apakah benul?", apakahBenul)
	// Output yang dihasilkan Printf tergantung flag yang digunakan, untuk melihat tipe data gunakan %T, hal ini disebut Go verbs.
	fmt.Printf("dengan tipe data %T\n", apakahBenul)

	// * Variable without data type (short declaration)
	name3 := "Kira Yoshikage"
	age2 := "33"

	fmt.Println("Nama saya", name3, "Umur saya", age2)

	// * Multiple variable declarations
	var student1, student2, student3 string = "s1", "s2", "s3"
	fmt.Println(student1, student2, student3)

	var name4, age3, address = "Nopal", 21, "Jalan diponegoro"
	// Coba printf lagi
	fmt.Printf("Halo saya %s, umur saya %d, rumah saya %s", name4, age3, address)

	/*
		Underscore variable, pada golang tidak boleh ada pendeklarasian variabel yang tidak dipakai
		untuk menanggulangi error ini dapat menginisialisasi variabel yang tidak dipakai tersebut pada
		underscore variabel seperti di bawah
	*/
	var name5, age4 = "Naldy", 23
	_, _ = name5, age4
}
