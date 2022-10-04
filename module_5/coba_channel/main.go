package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	c := make(chan string)
	go func() {
		c <- "Coba kirim channel"
	}()

	go func() {
		c <- "Coba kirim channel 2"
	}()

	go func() {
		c <- "Coba kirim channel 3"
	}()

	msg1 := <-c
	msg2 := <-c
	msg3 := <-c

	// pengiriman akan dijalankan secara asinkronus
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)

}
