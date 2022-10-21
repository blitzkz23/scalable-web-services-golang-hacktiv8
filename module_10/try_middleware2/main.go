package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", middleware(http.HandlerFunc(endpoint)))

	http.ListenAndServe(":8080", nil)
}

func endpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware pertama")

		next.ServeHTTP(w, r)
	})
}
