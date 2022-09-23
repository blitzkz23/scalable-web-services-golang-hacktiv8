package main

import (
	"fmt"
	"math"
	"strings"
)

// ! Function
func main() {
	fmt.Println()
	greet("Aldy", 21)
	fmt.Println()
	greet2("Aldy", "Jakarta")
	fmt.Println()
	fmt.Println(greetWithReturn("Hello", []string{"Aldy", "Budi", "Caca"}))

	fmt.Println(strings.Repeat("#", 50))
	var diameter float64 = 15

	var area, circumference = calculate(diameter)
	fmt.Printf("Luas lingkaran dengan diameter %.2f adalah %.2f dengan keliling %.2f\n", diameter, area, circumference)

	fmt.Println(strings.Repeat("#", 50))
	fmt.Println("Predefined return function")
	var area2, circumference2 = calculate2(diameter)
	fmt.Printf("Luas lingkaran dengan diameter %.2f adalah %.2f dengan keliling %.2f\n", diameter, area2, circumference2)

	fmt.Println(strings.Repeat("#", 50))
	fmt.Println("Variadic function")
	fmt.Println(printVariadic("Aldy", "Budi", "Caca"))
	fmt.Println("Variadic function 2")
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumVariadic(numberList...))
	fmt.Println("Variadic function 3")
	profileVariadic("Aldy", "Nasi Goreng", "Sate", "Bakso")
}

// * Normal function
func greet(name string, age int8) {
	fmt.Printf("Hello %s, you are %d years old", name, age)
}

// * Function with same parameter type
func greet2(name, address string) {
	fmt.Printf("Hello %s, you are from %s", name, address)
}

// * Function with return value
func greetWithReturn(msg string, names []string) string {
	var joinStr = strings.Join(names, ", ")
	// Sprintf untuk menampung hasil return value dan tiadk langsung di print
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// * Function with multiple return value
func calculate(d float64) (float64, float64) {
	// Menghitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	// Menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}

// * Function with predefined return value
func calculate2(d float64) (area, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

// * Variadic function, function yang memiliki parameter yang tidak terbatas
// tanda ... pada parameter menandakan fungsi ini adalah variadic
// parameter string yang diterima akan dikonversi menjadi slice yang mengandung tipe data map sehingga bisa diloop untuk mendapatkan key dan valuenya
func printVariadic(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// * Contoh variadic 2
func sumVariadic(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

// * Contoh variadic 3 menggabungkan parameter biasa dan variadic
func profileVariadic(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Printf("My name is %s and my favorite foods are %s", name, mergeFavFoods)
}
