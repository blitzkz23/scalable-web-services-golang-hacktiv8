package main

import (
	"fmt"
	"log"
	"reflect"
)

// ! Reflect digunakan untuk melakukan inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkanmemanipulasinya.  Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.

func main() {
	// * Identifikasi tipe data dari variabel
	var number = 21
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe variabel:", reflectValue.Type())

	// * Reflect dari struct
	type Person struct {
		Name string `required:"true"`
		Age  int    `required:"true"`
	}
	var p Person = Person{
		Name: "John",
	}

	var r reflect.Type = reflect.TypeOf(p)

	fmt.Println("Tipe variabel:", r)
	fmt.Println("Nama variabel:", r.Name())
	fmt.Println("Jumlah field:", r.NumField())

	// * Pengecekan jenis reflect menggunakan Kind() https://pkg.go.dev/reflect?utm_source=godoc#Kind
	if reflectValue.Kind() == reflect.Int {
		// Untuk mengakses nilai reflect dapat dengan pengecekan tipe seperti di atas, atau menggunakan metode interface seperti di bawah
		fmt.Println("Nilai variabel:", reflectValue.Int())
	}

	// * Mengakses nilai menggunakan interface method
	var nilai = reflectValue.Interface().(int)
	fmt.Println("Nilai variabel:", nilai)

	// * Validation menggunakan reflect
	validStruct(p)
}

// * Validation menggunakan reflect
func validStruct(data interface{}) {
	var reflectType reflect.Type = reflect.TypeOf(data)

	if reflectType.Kind() != reflect.Struct {
		log.Fatal("Payload harus berupa struct")
	}

	// Karena kita ingin melooping numfield dimana hanya dimiliki oleh tipe data struct
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)

		if field.Tag.Get("required") == "true" {
			reflectValuew := reflect.ValueOf(data).Field(i).Interface()

			if reflectValuew == "" {
				message := fmt.Sprintf("The %s field is required", field.Name)
				log.Fatal(message)
			}
		}
	}
	fmt.Println("Tipe valid")
}
