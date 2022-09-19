package main

import (
	"fmt"
)

func main() {
	// ! Slice merupakan tipe data yang mirip array namun tidak memiliki fixed-length, dan merupakan kategori reference type

	// * Slice ver.1
	var fruits = []string{"Melon", "Lemon", "Grape", "Cherry"}
	fmt.Printf("%#v\n", fruits)

	// * Slice ver. 2 with make function
	var foods = make([]string, 3)
	// verb %#v memberikan print dari representasi nilai variabel yang diformat
	fmt.Printf("%#v\n", foods)

	// * Slice append function
	foods = append(foods, "Kebab", "Salmon", "Takoyaki")

	fmt.Printf("%#v\n", foods)

	// * Append function with ellipsis, to append another slice to desired slice
	foods = append(foods, fruits...)
	fmt.Printf("%#v\n", foods)

	// * Slice copy function, if used replace the current slice with the copied slice, the first 4 index of food is replaced by the fruits based on the output
	println()
	fmt.Println("============================")
	fmt.Println("Foods before copy =>", foods)
	// newSlice merupakan hasil return dari jumlah element yang tercopy
	newSlice := copy(foods, fruits)

	fmt.Println("Fruits =>", fruits)
	fmt.Println("Foods after copy =>", foods)
	fmt.Println("Copied elements", newSlice)

	// * Teknik slicing yang mirip di python dengan cara menentukan index [start:stop]
	println()
	println("================================================")
	buahMurni := []string{"Anggur", "Apel", "Mangga", "Ceri", "Tomat"}
	fmt.Println("Original slice =>", buahMurni)

	// Slice dari index 1 hingga 3
	sliceOfBuah1 := buahMurni[1:4]
	fmt.Println("Slice buah 1 =>", sliceOfBuah1)

	// Slice dari index 0 hingga akhir
	sliceOfBuah2 := buahMurni[1:]
	fmt.Println("Slice buah 2 =>", sliceOfBuah2)

	// Slice dari awal hingga index 2
	sliceOfBuah3 := buahMurni[:3]
	fmt.Println("Slice buah 3 =>", sliceOfBuah3)

	// Slice dari awal hingga akhir
	sliceOfBuah4 := buahMurni[:]
	fmt.Println("Slice buah 4 =>", sliceOfBuah4)

	// * Combine slicing and append
	sliceOfBuah5 := append(sliceOfBuah4[3:], "Rambutan")
	fmt.Println("Slice buah 5 =>", sliceOfBuah5)

	// * Backing array membuat slice asal dapat terganti apabila slice tiruanya diubah, itu karena masih berada dalam backing array yang sama, contoh:
	println("===============================================")
	cobaBackingArray := []string{"naufal", "aldy", "pradana"}
	backingSlice := cobaBackingArray[1:]

	// Apabila ada yang diganti
	backingSlice[0] = "heru"

	// backingSlice index-0 aldy akan terganti heru dan cobaBackingArray index-1 aldy juga ikut terganti
	fmt.Println("Slice 1 =>", cobaBackingArray)
	fmt.Println("Slice 2 =>", backingSlice)

	// * Slice (cap function, digunakan untuk menghitung kapasitas slice)
	cobaCapSlice := []string{"Naufal", "Aldy", "Pradana", "Ganteng"}
	fmt.Println("cobaCapSlice cap:", cap(cobaCapSlice)) //4
	fmt.Println("cobaCapSlice len:", len(cobaCapSlice)) //4

	sliceOfCap1 := cobaCapSlice[0:3]
	fmt.Println("sliceOfCap1 cap:", cap(sliceOfCap1)) //4
	fmt.Println("sliceOfCap2 len:", len(sliceOfCap1)) //3

	sliceOfCap2 := cobaCapSlice[1:]
	fmt.Println("sliceOfCap2 cap:", cap(sliceOfCap2)) //3
	fmt.Println("sliceOfCap2 len:", len(sliceOfCap2)) //3

	// Kenapa sliceOfCap2 memiliki kapasitas lebih sedikit dari sliceOfCap1? dapat diilustrasikan sebagai berikut hasil sliceOfCap1 = ["Naufal", "Aldy", "Pradana", ...]
	// sliceOfCap2 = [..., "Aldy", "Pradana", "Ganteng"] karena sliceOfCap1 mengambil dari index 1, sehingga masih tersisa 1 kapasitas terakhir pada index terakhir, sedangkan pada sliceOfCap2 slice dimulai hitung dari index ke x pada [x:y] sehingga menyebabkan kapasitas berubah.

	// * Create new backing array
	// ketika ingin mendapat element dari slice yang sudah ada dengan backing array baru maka gunakan append

	cars := []string{"Ford", "Tesla", "Mustang", "Tiger"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	// Setelah index 0 diganti tidak mempengaruhi newCars
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

}
