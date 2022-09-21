package main

import (
	"fmt"
)

func main() {
	// ! If else
	var currentYear = 2022

	// this are tempprary variable, that could only be accessed in certain scope, dimana pada kode di bawah age < 17 merupakan sebuah kondisi, dan age sendiri menghasilkan angka 2024 yang akan mentrigger block elses
	if age := currentYear - 1998; age < 17 {
		fmt.Println("Anda belum boleh membuat SIM")
	} else {
		fmt.Println("Anda sudah boleh membuat SIM")
	}

	if bilangan := currentYear - 1991; bilangan%2 == 0 {
		fmt.Println("Sebuah bilangan genap")
	} else {
		fmt.Println("Sebuah bilangan ganjil")
	}

	// ! Switch
	var score = 4
	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Mayan")
		fallthrough // ini memungkinkan pengecekan case selanjutnya walaupun pada case ini telah terpenuhi
	case score < 5:
		fmt.Println("Mayan sih mayan tapi ya.. ditingkatkan lagi lh")
	default:
		{
			fmt.Println("Belajar lagi yh")
			fmt.Println("Nilai kok bgituh")
		}
	}

	// ! Nested conditions (if + switch)
	var score2 = 10

	if score2 > 7 {
		switch score {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Manteb")
		}
	} else {
		if score == 5 {
			fmt.Println("Not bad")
		} else if score == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if score == 0 {
				fmt.Println("Bro tf do you even try?")
			}
		}
	}

}
