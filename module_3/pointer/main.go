package main

import (
	"fmt"
	"strings"
)

// ! Pointer merupakan sebuah tipe data yang berisi alamat memory dari sebuah variable.  Contoh, ada variable x yang berisi nilai 10, maka pointer dari x adalah alamat memory dari x bukan nilai x itu sendiri.
func main() {
	fmt.Println()
	var firstNumber int = 4
	// secondNumber adalah pointer dari firstNumber yang menyimpan alamat memory dari firstNumber, tanda & untuk assign nilai pointer
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value)", firstNumber)
	fmt.Println("firstNumber (memory adress)", &firstNumber)

	// untuk mengakses nilai dari pointer, dimana secondNumber adalah pointer dari firstNumber, maka kita bisa mengakses nilai dari firstNumber dengan cara *secondNumber
	fmt.Println("secondNumber (value)", *secondNumber)
	// secondNumber di bawah berisi memori address yang sama dengan &firstNumber
	fmt.Println("secondNumber (memory adress)", secondNumber)

	// * Pointer menyimpan alamat dari suatu variabel, maka apabila nilai pada pointer tersebut diubah variabel awal juga akan ikut berubah, contoh:

	fmt.Println(strings.Repeat("#", 50))
	var firstPerson string = "Aldy"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value)", firstPerson)
	fmt.Println("firstPerson (memory adress)", &firstPerson)

	fmt.Println("secondPerson (value)", *secondPerson)
	fmt.Println("secondPerson (memory adress)", secondPerson)

	*secondPerson = "Naufal"

	fmt.Println("Setelah diganti nilai pointer:")
	fmt.Println("firstPerson (value)", firstPerson)
	fmt.Println("firstPerson (memory adress)", &firstPerson)

	fmt.Println("secondPerson (value)", *secondPerson)
	fmt.Println("secondPerson (memory adress)", secondPerson)

	// * Pointer as a parameter
	fmt.Println(strings.Repeat("#", 50))
	var number int = 4
	fmt.Println("number (before)", number)
	changeValue(&number)
	fmt.Println("number (after)", number)
}

func changeValue(number *int) {
	*number = 10
}
