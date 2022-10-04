package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

// ! net/http merupakan package yang digunakan untuk membuat aplikasi berbasis web seperti routing, templating, dll.
func main() {
	// * http.HandleFunc untuk keperluan routing yang menerima 2 parameter, params pertama routenya params kedua function yang memiliki parameter http.ResponseWriter dan http.Request.
	http.HandleFunc("/", greet)

	http.ListenAndServe(PORT, nil)
}

// * http.ResponseWriter merupakan sebuah interface dengan method-method yang digunakan untuk menulis response ke balik kepada client.  Sementara http.Request merupakan sebuah struct yang mendapatkan berbagai macam data request seperti form value, header, url params, dll.
func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"

	// * fmt.Fprint digunakan untuk mengirim response.
	fmt.Fprint(w, msg)
}
