package main

// func main(){
// greet("Airell", "Hobby", 24)

// var firstValue, secondValue, thirdValue  = generateWords("Lorem Ipsum Dolor Sit Amet", 100, map[string]string{"name" : "airell"})

// fmt.Println("firstValue =>", firstValue)
// fmt.Println("secondValue =>", secondValue)
// fmt.Println("thirdValue =>", thirdValue)

// generateVariadicFunction("Johan", true, false)
// type myfunc func(a int, b int) int

// var sum myfunc = func(a int, b int) int {
// 	return a + b
// }

// var divide func(a int, b int) int = sum

// var substract int = func(a int, b int) int {
// 	return a - b
// }(20, 10)

// fmt.Println(sum(10, 10))

// fmt.Println(substract)

// func(name string) {
// 	fmt.Println("Aku adalah function yang menginvoke diri sendiri")
// }("Airell")

// var checkOddNumber func (a int) bool = func (a int) bool {
// 	return a % 2 != 0
// }

// tryCallback(10, checkOddNumber)

// var countName func (name string) int  = countByte()

// var countName2 int = countByte()("Airell")

// fmt.Println(countName("Airell"))

// fmt.Println(countName2)
// }

// func checkOddNumber(a int) bool {
// 	return a % 2 != 0
// }

// func countByte() func (name string) int {
// 	return func(name string) int {
// 		return len(name)
// 	}
// }

// func tryCallback(num int ,isOddNumber func (a int) bool) {
// 	var result = isOddNumber(num)

// 	if result {
// 		fmt.Println("angka yang kamu berikan adalah angka ganil")
// 	}else {
// 		fmt.Println("angka yang kamu berikan adalah angka genap")
// 	}
// }

// func generateVariadicFunction (name string, conditionals ...bool) {
// 	fmt.Println("name =>",name)
// 	fmt.Println("conditionals =>", conditionals)
// }

// func generateWords(words string, randomNumber int, person map[string]string) (firstValue string, secondValue int, thirdValue map[string]string){
// 	firstValue = words
// 	secondValue = randomNumber
// 	thirdValue = person
// 	return
// }

// func greet(name string, hobby string, age int8) {
// 	fmt.Printf("Halo %s ini dari greet() \n", name)
// 	fmt.Printf("Kamu suka %s yah? \n", hobby)
// 	fmt.Printf("Umur kamu %d kan? \n", age)
// }