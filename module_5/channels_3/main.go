package main

import (
	"fmt"
	"time"
)

// ! Channel Select merupakan sebuah fitur bahasa go yang memungkinkan kita untuk menggunakan lebih dari satu channel untuk proses komunikasi antar goroutine.
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c2 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		// ! Kondisi case akan dijalankan jika channel c1 menerima data yang kemudiannya akan disimpan di variable msg1, begitu juga dengan case c2.
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
