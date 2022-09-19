package main

import "fmt"

func main() {
	// ! Loopings
	// * First way
	fmt.Println("Looping i")
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i+1)
	}

	// * Second way
	fmt.Println("Looping x")
	x := 0

	for x < 5 {
		fmt.Println("Angka", x)
		x++
	}

	// * Third way
	fmt.Println("Looping y")
	y := 0

	for {
		fmt.Println("Angka", y)

		y++
		if y == 5 {
			break
		}
	}

	// * Trying continue and break
	fmt.Println("Looping coba continue and break")
	for index := 0; index < 10; index++ {
		if index%2 == 1 {
			continue
		}

		if index > 8 {
			break
		}

		fmt.Println("Angka", index)
	}

	// * Nested Looping
	fmt.Println("Nested looping")
	for idx := 0; idx < 5; idx++ {
		for j := idx; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// * Pemberian label untuk menghentikan loop
	fmt.Println("Nested looping with label")
outerLoop: // ini labelnya
	for indexLuar := 0; indexLuar < 3; indexLuar++ {
		fmt.Println("Perulangan ke - ", indexLuar+1)
		for indexDalam := 0; indexDalam < 3; indexDalam++ {
			if indexLuar == 2 {
				break outerLoop
			}
			fmt.Print(indexDalam, " ")
		}
		fmt.Println()
	}
}
