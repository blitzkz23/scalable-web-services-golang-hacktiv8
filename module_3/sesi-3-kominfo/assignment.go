package main

import (
	"fmt"
	"os"
	_ "strconv"
)

type Person struct {
	Id string
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}


func main() {
	var people []Person = []Person{
		{
			Id: "1",
			Nama: "Airell",
			Alamat: "Jakarta",
			Pekerjaan: "Ngoding dan sebat",
			Alasan: "Ya karena hobi",
		},
		{
			Id: "2",
			Nama: "Andi",
			Alamat: "Bekasi",
			Pekerjaan: "Software Engineer (BE)",
			Alasan: "Ya karena suka",
		},
	}


	var args = os.Args

	fmt.Println(getPersonById(args[1], people))

}

func getPersonById(id string, people []Person) string {

	var conditional bool = false

	var peopleIndex = 0

	for index, value := range people {
		if(value.Id == id) {
			conditional = true
			peopleIndex = index
			break
		}
	}

	if conditional {
		return fmt.Sprintf("Data person ditemukan => %+v", people[peopleIndex])
	}else {
		return "Data person tidak ditemukan"
	}
}