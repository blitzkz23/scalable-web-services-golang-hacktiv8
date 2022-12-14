package main

import "fmt"

// ! Empty Interface merupakan tipe data yang menerima segala tipe data, namun tipe data tersebut tidak dapat dicantumkan secara eksplisit sehingga ambigu.
func main() {
	var randomValue interface{}
	_ = randomValue

	// * Random Value bisa dimasuki tipe data apa saja
	randomValue = "Jalan Sudirman"
	randomValue = 20
	randomValue = true
	randomValue = []string{"Airell", "Ari", "Ariell"}
	fmt.Println(randomValue)

	var v interface{}
	v = 20

	// v = v * 9  // ! ini akan error karena v adalah interface kosong yang diisi tipe data int dan bukan int itu sendiri sehingga tidak bisa dilakukan operasi matematika

	// * Untuk menanggulangi hal tersebut, kita bisa melakukan type assertion.  Joka tipe data interface ok adalah true maka kita bisa melakukan operasi matematika.
	if value, ok := v.(int); ok {
		v = value * 9
		fmt.Println(v)
	}

	// * Map & Slice with empty interface
	// Slice of empty interface yang bisa diisi tipe data apa saja
	rs := []interface{}{1, "Naufal", true, 2, "Airell", false}

	// Map dengan key berupa string, dan value berupa interface kosong yang bisa diisi tipe data apa saja
	rm := map[string]interface{}{
		"nama":   "Naufal",
		"umur":   20,
		"status": true,
	}

	_, _ = rs, rm

	// * Type switching
	var sliceData []interface{} = []interface{}{
		"Naufal",
		20,
		true,
	}

	for index, v := range sliceData {
		switch value := v.(type) {
		case int:
			fmt.Printf("Data %d merupakan sebuah %T dengan nilai %d \n", index, value, value)
		case string:
			fmt.Printf("Data %d merupakan sebuah %T dengan nilai %s \n", index, value, value)
		case bool:
			fmt.Printf("Data %d merupakan sebuah %T dengan nilai %t \n", index, value, value)
		default:
			fmt.Println("Tipe data tidak diketahui")
		}
	}

	// * Interface function with variadic parameter
	receiveRandomData("Naufal", 20, true, 3.14, []string{"Airell", "Ari", "Ariell"})
}

// * Interface function with variadic parameter
func receiveRandomData(a ...interface{}) {
	fmt.Println(a...)
}
