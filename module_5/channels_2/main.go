package main

import (
	"fmt"
	"strings"
	"time"
)

// ! Pada dasarnya channel bersifat unbuffered/tidak buffer di memori dimana proses penerimaan dan pengirim bersifat asinkronus.  Namun untuk buffered channel kita bisa menentukan kapasitas bufernya, dan selama jumlah data yang dikirim melalui unbuffered channel tidak melebihi kapasitasnya, maka proses pengiriman data akan bersifat asynchronous.
// ! Proses pengiriman dan penerimaan pada unbuffered bersifat sinkronus/tunggu menunggu.  Semisal goroutine ke-10 mengirim data ke goroutine main setelah dikirim dan diterima, baru goroutine lain bisa mengirim data.
func main() {
	fmt.Println()

	// * Unbuffered channel pengiriman blocking, penerimaan blocking
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		// Pengiriman
		c <- 10
		// ! Karena channel bersifat unbuffered, maka proses di bawah pengiriman data akan berhenti sampai data diterima oleh goroutine yang menerima data dari channel.
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	// Penerimaan
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	close(c1)
	time.Sleep(time.Second)

	// * Buffered channel
	fmt.Println(strings.Repeat("#", 50))
	c2 := make(chan int, 3)
	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d sends data into the channel\n", i)
			// Pengiriman asinkronus selama slot masih ada
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}

		close(c)
	}(c2)

	fmt.Printf("main goroutine sleeps for 2 seconds\n")
	time.Sleep(time.Second * 2)

	// * Menerima data dari channel dengan loop
	// Penerimaan sinkronus
	for v := range c2 {
		fmt.Println("main goroutine received data from channel:", v)
	}

	// ! Berdasarkan hasil dari print, cara kerja buffered render yaitu pengiriman data akan dieksekusi hingga kapasitas buffer terpenuhi lalu proses penerimaan data oleh main akan dieksekusi, baru proses pengiriman data akan dilanjutkan kembali jika masih ada data yang dikirim.

}
