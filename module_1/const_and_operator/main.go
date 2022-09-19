package main

import "fmt"

func main() {

	// ! Constant
	// Apabila membuat constant harus langsung dideklarasikan valuenya
	const base_url string = "https://nyobaconst.com"
	const phi_number = 3.14

	fmt.Printf("sebuah api call %s/%.2f \n", base_url, phi_number)

	// ! Operators
	// Ada Math operator, logic operator, relational operator

	// * Math Op
	fmt.Printf("sebuah nilai %d \n", 2+3)

	// * Logic Op
	var benul = true
	var salah = false

	fmt.Printf("Benul dan salah (%t) \n", benul && salah)
	fmt.Printf("Benul atau salah (%t) \n", benul || salah)
	fmt.Printf("Tidak benul (%t) \n", !benul)

	// * Relational Op
	var nama1 = "Aldy"
	var nama2 = "aldy"
	fmt.Printf("Sama? %t", nama1 == nama2)
}
