package main

import (
	"fmt"
	"strings"
)

func main() {
	// ! Array pada go memiliki fixed-length dan harus ditentukan dari awal
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var myNameArray = [3]string{"Naufal", "Aldy", "Pradana"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", myNameArray)

	// * Loop through element
	var buah = [3]string{"Kiwi", "Anggur", "Cherry"}

	// First way
	for index, value := range buah {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	fmt.Println(strings.Repeat("#", 25))

	// Second way
	for i := 0; i < len(buah); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, buah[i])
	}

	// * Multidimensional Array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		fmt.Println("Ini nilai", arr)
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
