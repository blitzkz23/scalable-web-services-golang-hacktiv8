package main

import "fmt"

// ! Struct ini mirip data class di kotlin lah yaw

type Address struct {
	city     string
	province string
}

type Employee struct {
	// Ini adalah embedded struct
	name     string
	salary   int
	division string
	address  Address
}

// * Anonymous struct
type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println()

	// * Memberi value pada struct
	var employee1 Employee

	employee1.name = "Aldy"
	employee1.salary = 100000000
	employee1.division = "IT"

	// embedded property
	employee1.address.city = "Jakarta"
	employee1.address.province = "DKI Jakarta"

	fmt.Printf("employee1: %v\n", employee1)

	// * Initialize struct secara langsung
	var employee2 = Employee{name: "Budi", salary: 3000000, division: "IT"}
	employee2.address.city = "Jakarta"
	employee2.address.province = "DKI Jakarta"
	fmt.Printf("employee2: %v\n", employee2)

	// * Menggunakan pointer untuk mengubah value struct
	var employee3 *Employee = &employee2
	employee3.name = "Caca"
	fmt.Println("Nilai nama employee2 setelah diubah oleh employee3 :\n", employee2.name)
	fmt.Println("Nilai nama employee3 :\n", employee3.name)

	// * Inisialisasi anonymous strict
	var soldier = struct {
		person   Person
		division string
		rank     string
	}{}
	soldier.person = Person{name: "Aldy", age: 21}
	soldier.division = "Bomber"
	soldier.rank = "Lieutenant"
	fmt.Printf("soldier: %v\n", soldier)

	// * Slice of struct
	var people = []Person{
		{name: "Aldy", age: 21},
		{name: "Budi", age: 22},
		{name: "Caca", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	// * Slice of anonymous struct
	var warrior = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Aldy", age: 21}, division: "Bomber"},
		{person: Person{name: "Budi", age: 22}, division: "Bomber"},
		{person: Person{name: "Caca", age: 23}, division: "Bomber"},
	}

	for _, v := range warrior {
		fmt.Printf("%+v\n", v)
	}
}
