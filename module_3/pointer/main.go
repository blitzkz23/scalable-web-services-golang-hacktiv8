package main

import (
	"fmt"
	"strings"
)

// ! Pointer merupakan sebuah tipe data yang berisi alamat memory dari sebuah variable.  Contoh, ada variable x yang berisi nilai 10, maka pointer dari x adalah alamat memory dari x bukan nilai x itu sendiri.
// Pointer banyak digunakan nantti saat membangun API karena saat data tidak ada dia bisa mereturn nilai nil
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

	// Untuk pemanggilan nilai dari pointer, kita bisa menggunakan *secondPerson yang bisa dinamakan dengan dereferencing
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

	// * Function that return pointer
	fmt.Println(strings.Repeat("#", 50))
	var person *string = getPerson("Aldy")
	fmt.Println("person mem. address found =", person)
	fmt.Println("person value found =", *person)
	// Coba cari person yang tidak ada
	var person2 *string = getPerson("Rizkt")
	// Kalau tidak menggunakan condition akan mengalami panic error saat dereferencing person2
	if person2 != nil {
		fmt.Println("person2 mem. address found =", person2)
		fmt.Println("person2 value found =", *person2)
	} else {
		fmt.Println("person2 not found")
	}

	// * Function that change value based on pointer
	var name1 string = "Johnson"
	changeValueName(&name1)
	fmt.Println("name1 =", name1)
}

func changeValue(number *int) {
	*number = 10
}

// * func yang return pointer dari suatu string
func getPerson(name string) *string {
	var students []string = []string{"Aldy", "Naufal", "Rizky"}

	for _, student := range students {
		if student == name {
			return &student
		}
	}

	// Kembalikan nilai nil jika tidak ada
	return nil
}

// * func yang change value berdasar pointer
func changeValueName(name *string) {
	*name = "Anon"
}
