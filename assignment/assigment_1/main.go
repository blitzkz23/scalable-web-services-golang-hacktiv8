package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct for student
type student struct {
	id                   int
	name                 string
	address              string
	job                  string
	reasonToChooseGolang string
}

var listOfStudent []student

func main() {
	listOfStudent = []student{
		{id: 1, name: "Naufal", address: "Semarang", job: "Programmer", reasonToChooseGolang: "Easy to learn"},
		{id: 2, name: "Rizky", address: "Jakarta", job: "Programmer", reasonToChooseGolang: "Microservices friendly"},
		{id: 3, name: "Rahmat", address: "Bandung", job: "Programmer", reasonToChooseGolang: "I love suffering"},
		{id: 4, name: "Rizal", address: "Surabaya", job: "Programmer", reasonToChooseGolang: "I like challenge"},
		{id: 5, name: "Rifqi", address: "Malang", job: "Programmer", reasonToChooseGolang: "Just because"},
		{id: 6, name: "Rizki", address: "Bali", job: "Programmer", reasonToChooseGolang: "I like it"},
		{id: 7, name: "Rahmat", address: "Bogor", job: "Programmer", reasonToChooseGolang: "I like it"},
		{id: 8, name: "Rizal", address: "Bekasi", job: "Programmer", reasonToChooseGolang: "I like it"},
		{id: 9, name: "Rifqi", address: "Depok", job: "Programmer", reasonToChooseGolang: "I like it"},
		{id: 10, name: "Rizki", address: "Tangerang", job: "Programmer", reasonToChooseGolang: "I like it"},
	}

	// * Get student by id
	id, _ := strconv.Atoi(os.Args[1])
	var searchedStudent *student = getStudentById(&id)

	if searchedStudent != nil {
		printStudentInfo(searchedStudent)
	} else {
		fmt.Println("Murid dengan id", id, "tidak ditemukan")
	}
}

// * Get student by id @params pointer of id, @return pointer of student
func getStudentById(id *int) *student {
	for _, student := range listOfStudent {
		if student.id == *id {
			// kembalikan pointer dari student
			return &student
		}
	}
	return nil
}

// * Print student info
func printStudentInfo(student *student) {
	fmt.Println(strings.Repeat("#", 25))
	fmt.Println("Murid dengan id", student.id, ":")
	fmt.Println("Nama:", student.name)
	fmt.Println("Alamat:", student.address)
	fmt.Println("Pekerjaan:", student.job)
	fmt.Println("Alasan memilih golang:", student.reasonToChooseGolang)
	fmt.Println(strings.Repeat("#", 25))
}
