package main

import (
	"fmt"
	"runtime"
	"time"
)

// ! Concurrency adalah mengeksekusi sebuah proses secara independen atau berhadapan dengan beberapa tugas dalam waktu yang sama.
// ! Goroutine merupakan sebuah thread ringan pada bahasa Go untuk melakukan proses secara concurrent.  1 goroutine hanya membutuhkan 2kb memory saja, sedangkan 1 thread membutuhkan 1-2mb memory.
func main() {
	fmt.Println("Main Execution Started")
	// pemanggilan goroutine diawali keyword go
	go firstProcess(8)
	secondProcess(8)

	// function main juga merupakan goroutine sehingga num of goroutine akan menjadi 2, dan main tidak menunggu eksekusi goroutine lainya sehingga hasil dari firstProcess tidak akan ditampilkan
	fmt.Println("No. of Goroutines", runtime.NumGoroutine())

	// karena eksekusi goroutine main ditahan selama 2 detik, maka hasil dari firstProcess akan ditampilkan
	time.Sleep(time.Second * 2)
	fmt.Println("Main Execution Finished")

	fmt.Println(runtime.NumCPU())

}

func firstProcess(index int) {
	fmt.Println("First Process started")
	for i := 1; i <= index; i++ {
		fmt.Println("i = ", i)
	}
	fmt.Println("First Process finished")
}

func secondProcess(index int) {
	fmt.Println("Second Process started")
	for j := 1; j <= index; j++ {
		fmt.Println("j = ", j)
	}
	fmt.Println("Second Process finished")
}
