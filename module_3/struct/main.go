package main

import "fmt"

// ! Struct ini mirip data class di kotlin lah yaw

// * Apabila dimulai dengan huruf besar maka struct tersebut bisa diakses dari luar package, jika kecil maka private.  Ini berlaku juga untuk property struct.
type address struct {
	city     string
	province string
}

type Employee struct {
	// Ini adalah embedded struct
	Name     string
	salary   int
	division string
	address  address
}

// * Anonymous struct
type Person struct {
	name string
	age  int
}

// * Struct dalam struct
type BankAccount struct {
	TotalSavings int
	VaNumber     int
	Owner        Person
	// Atau bisa dibuat
	// Owner struct {
	// 	Name string
	// 	Age  int
	// }
}

// * Juga bisa menggunakan pointer untuk mengubah value struct
func (p *Person) changeName(newName string) {
	p.name = newName
}

// * Coba method untuk struct
func (p Person) sayHello() {
	fmt.Println("Hello", p.name)
}

// * Membuat function yang melekat pada struct, sehingga ini dapat dikatakan method milik employee.
func (e Employee) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", e.Name)
	// Kita juga bisa mengambil detail dari instance yang memanggil function ini
	fmt.Printf("%+v\n", e)
}

func main() {
	fmt.Println()

	// * Memberi value pada struct
	var employee1 Employee

	employee1.Name = "Aldy"
	employee1.salary = 100000000
	employee1.division = "IT"
	employee1.sayHello("Naufal")

	// embedded property
	employee1.address.city = "Jakarta"
	employee1.address.province = "DKI Jakarta"

	fmt.Printf("employee1: %v\n", employee1)

	// * Initialize struct secara langsung
	var employee2 = Employee{Name: "Budi", salary: 3000000, division: "IT"}
	employee2.address.city = "Jakarta"
	employee2.address.province = "DKI Jakarta"
	fmt.Printf("employee2: %v\n", employee2)

	// * Menggunakan pointer untuk mengubah value struct
	var employee3 *Employee = &employee2
	employee3.Name = "Caca"
	fmt.Println("Nilai nama employee2 setelah diubah oleh employee3 :\n", employee2.Name)
	fmt.Println("Nilai nama employee3 :\n", employee3.Name)

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

	// * Coba buat struct dengan property yang berisi struct
	var person1 = Person{name: "Aldy", age: 21}

	var bankAccount BankAccount = BankAccount{
		TotalSavings: 100000000000,
		VaNumber:     123456789,
		Owner:        person1,
	}
	fmt.Printf("%+v\n", bankAccount)
	fmt.Println("Nama pemilik rekening bank :", bankAccount.Owner.name)
	fmt.Println("Umur pemilik rekening bank :", bankAccount.Owner.age)
	fmt.Println("Total tabungan :", bankAccount.TotalSavings)
	fmt.Println("Nomor rekening :", bankAccount.VaNumber)
	// Juga bisa invoke function yang ada di struct person
	bankAccount.Owner.sayHello()

	// * Coba ubah value struct person
	bankAccount.Owner.changeName("Naufal")
	fmt.Println("Nama pemilik rekening bank :", bankAccount.Owner.name)

	// * Coba pointer of struct, pointer dari struct tidak perlu didereference
	var bankAccount2 *BankAccount = &BankAccount{
		TotalSavings: 100000000000,
		VaNumber:     123456789,
		Owner:        person1,
	}
	fmt.Printf("%+v\n", bankAccount2)
}
