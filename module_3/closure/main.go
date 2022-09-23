package main

import (
	"fmt"
	"strings"
)

// ! Closure merupakan sebuah anonymous function yang bisa disimpan dalam sebuah variable/dijadikan nilai return dari sebuah function.
func main() {
	fmt.Println()
	// * Closure ver. 1
	// Buath sebuah variabel untuk menerima bilangan genap
	var evenNumbers = func(numbers ...int) []int {
		// Create empty slice
		var result []int
		// Append number to result if number is even
		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(evenNumbers(numbers...))
	fmt.Println(strings.Repeat("#", 50))

	// * Closure IIFE (Immediately Invoked Function Expression)
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
		// di bawah merupakan pemanggilan IIFE
	}("katak")

	fmt.Println(isPalindrome)

	// * Closure as return value
	fmt.Println(strings.Repeat("#", 50))
	var students = []string{"Aldy", "Budi", "Caca", "Dedi", "Euis"}
	var findStudent = findStudent(students)
	// find adalah closure yang direturn function findstudent, dimana Ia menerima parameter string dan mereturn string
	fmt.Println(findStudent("Aldy"))

	// * Closure (Callback), merupakan closure yang dijadikan sebagai parameter sebuah function
	fmt.Println(strings.Repeat("#", 50))
	var numbers2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var findOdd = findOddNumber(numbers2, func(number int) bool {
		// Isi dari callback function didefinisikan disini
		return number%2 != 0
	})
	fmt.Println("Total odd numbers:", findOdd)

}

// * Closure as return value
// function findStudent akan menerima parameter slice of string, dan mereturn sebuah closure berupa function, yang menerima parameter string dan mereturn string
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s does'nt exist!", s)
		}
		return fmt.Sprintf("%s is at  %d", student, position+1)
	}
}

// * Closure (Callback)
// func findOddNumber(numbers []int, callback func(int) bool) int {

// Mempersingkat pemanggilan callback dengan type alias
type isOddNum func(int) bool

func findOddNumber(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		// Oke disini pemanggilan callback
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
