package main

import (
	"fmt"
	"strings"
)

// ! Channel merupakan sebuah mekanisme goroutine untuk berkomunikasi satu sama lain.
// ! syntax:
// ! c chan string // nama channel
// ! c <- result // mengirim data ke channel
// ! <-c // menerima data dari channel

func main() {
	fmt.Println()
	// * Membuat channel dengan tipe data string
	c := make(chan string)

	go introduce("Naufal", c)
	go introduce("Rizky", c)
	go introduce("Rahmat", c)

	// * Function main menerima data dari channel, yang masing2 disimplan oleh msg1, msg2, msg3.  Hasil yang diprint pun tidak menentu untuk urutanya. ini terjadi karena goroutine memiliki sifat asinkronus dan tidak saling menunggu.
	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	// ! Akan error karena semua gorooutine sudah selesai, dan channel tidak memiliki data lagi
	// msg4 := <-c
	// fmt.Println(msg4)

	// ! Maka channel harus diclose
	close(c)

	// ! Karena function main juga merupakan sebuah goroutine, maka syntax di atas merupakan contoh proses komunikasi antar goroutine main dengan goroutine introduce.

	// * Chanel with anonymous function
	fmt.Println(strings.Repeat("#", 50))
	c2 := make(chan string)

	students := []string{"Naufal", "Rizky", "Rahmat"}

	for _, value := range students {
		// goroutine function with channel
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hello, my name is %s", student)
			c2 <- result
		}(value)
	}

	for i := 1; i <= 3; i++ {
		print(c2)
	}

	close(c2)
	fmt.Println(strings.Repeat("#", 50))

}

// Fungsi yang mengirim data ke channel main
func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hello, my name is %s", student)

	// * Mengirim data ke channel
	c <- result
}

// Fungsi yang menerima data dari channel main
func print(c2 chan string) {
	// * Menerima data dari channel
	fmt.Println(<-c2)
}
