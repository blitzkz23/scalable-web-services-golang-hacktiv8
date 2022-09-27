package main

import (
	"fmt"
	"reflect"
)

// ! Reflect digunakan untuk melakukan inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkanmemanipulasinya.  Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.

func main() {
	// * Identifikasi tipe data dari variabel
	var number = 21
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe variabel:", reflectValue.Type())

	// Pengecekan jenis reflect menggunakan Kind() https://pkg.go.dev/reflect?utm_source=godoc#Kind
	if reflectValue.Kind() == reflect.Int {
		// Untuk mengakses nilai reflect dapat dengan pengecekan tipe seperti di atas, atau menggunakan metode interface seperti di bawah
		fmt.Println("Nilai variabel:", reflectValue.Int())
	}

	// * Mengakses nilai menggunakan interface method
	var nilai = reflectValue.Interface().(int)
	fmt.Println("Nilai variabel:", nilai)

	// * Mengidentifikasi informasi method
	// func (s *student) SetName(name string) {
	// 	s.Name = name
	// }
	// var s1 = &student{name: "Airell", age: 20}
	// fmt.Println("Nama:", s1.name)

	// var reflectValue = reflect.ValueOf(s1)
	// var method = reflectValue.MethodByName("SetName")
	// method.Call([]reflect.Value{
	// 	reflect.ValueOf("Ari"),
	// })

	// fmt.Println("Nama:", s1.name)

}
