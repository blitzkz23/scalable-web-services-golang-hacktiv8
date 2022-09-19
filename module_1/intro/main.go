package main

import "fmt"

// Apabila sebuah variabel ada pada scope global seperti ini, tidak akan membuat error apabila tidak dipakai.
var appName = "Main App"

func main() {
	// Sebuah comment
	/*
		Sebuah comment panjang
	*/
	fmt.Println("Hello World", "Bisa dipisah2", "gini", "anjir paramsnya")
	fmt.Print("Yang ini tanpa", "spasi ygy")
	fmt.Println("Bukti")
	helloWorld()
}
