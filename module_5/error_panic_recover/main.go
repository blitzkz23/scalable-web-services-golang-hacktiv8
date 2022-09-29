package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// ! Error merupakan sebuah struct yang digunakan untuk menangkap error yang terjadi pada program.  Umumnya, nilai error di-return pada posisi terakhir dari sebuah function.

	var number int
	var err error

	// Value 123gh adalah string yang tidak dapat diubah menjadi integer, sehingga akan menghasilkan error yang dikembalikan oleh strconv.Atoi
	number, err = strconv.Atoi("123GH")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// * Custom error

	// * Apabila fungsi recover jalan saat ada error pada goroutine (main), pesan panic error tidak ditampilkan
	defer cathErr()

	var password string
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		// * Panic digunakan untuk menampilkan stack trace error sekaligus menghentikan flow goroutine (main juga goroutine).  Setelah ada panic, proses akan terhenti, apapun setelah tidakdi-eksekusi kecuali proses yang sudah di-defersebelumnya (akan muncul sebelum panic error).
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

// * Validation with custom error
func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "Valid password", nil
}

// * Function recover digunakan untuk menangkaperror yang terjadi pada sebuah Goroutine.
func cathErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running smoothly")
	}
}
