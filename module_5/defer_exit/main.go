package main

import (
	"fmt"
	"os"
	"strings"
)

// ! Defer merupakan sebuah keyboard pada bahasa go yang digunakan untuk memanggil sebuah yang dimana proses eksekusi akan ditahan hingga block sebuah fungsi selesai dieksekusi.
func main() {
	// * Defer 1
	fmt.Println(strings.Repeat("#", 50))
	defer fmt.Println("defer function 1 starts to execute") // fungsi ini akan ditahan sampe akhir
	fmt.Println("Hai everyone")
	fmt.Println("Wecome to Go Programming Language")

	// * Defer 2 - Call to defer func
	fmt.Println(strings.Repeat("#", 50))
	callDeferFunc()
	fmt.Println("Hai everyone")

	// * Defer 3 - Dengan exit
	fmt.Println(strings.Repeat("#", 50))
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
