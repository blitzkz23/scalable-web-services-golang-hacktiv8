package main

import (
	"fmt"
	"sync"
)

// ! Waitgroup merupakan sebuah struct dari package sync, yang digunakan untuk melakukan sinkronisasi terhadap goroutine.

func main() {
	fmt.Println()
	fruits := []string{"apple", "banana", "grape", "orange"}
	var waitgroup sync.WaitGroup

	for index, fruit := range fruits {
		// Menambah conuter waitgroup sebanyak jumlah goroutine yang akan dijalankan
		waitgroup.Add(1)
		go printFruit(index, fruit, &waitgroup)
	}

	waitgroup.Wait()
}

// Argumen waitfroup di bawah merupakan pointer dari waitgroup agar method mengandung memori yang sama
func printFruit(index int, fruit string, waitGroup *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	waitGroup.Done()
}
