package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "John", Age: 30, Division: "IT"},
	{ID: 2, Name: "Doe", Age: 25, Division: "HR"},
	{ID: 3, Name: "Jane", Age: 27, Division: "IT"},
}

var PORT = ":8080"

// ! go run api.go untuk menjalankan
func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employees/create", createEmployee)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

// * Coba method GET
func getEmployees(w http.ResponseWriter, r *http.Request) {
	// * Method Header dari interface http.ResponseWriter yang dichain dengan method Set.  Hal ini dilakukan untuk menentukan bentuk dari data response yang akan dikirimkan ke client.
	// w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// * Method Encode dari package json yang digunakan untuk mengubah data struct employees menjadi bentuk json.
		// json.NewEncoder(w).Encode(employees)
		// return

		// * Parse data menjadi halaman html, dan kalau suatu function bisa mengembalikan error harus ditangkap.
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		return
	}

	// * Jika method yang dikirimkan oleh client bukan method GET, maka akan dikirimkan response error status 400 bad request.
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	// * Untuk mengembalikan data berupa json
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		// * Nilai input didapatkan dari method FormValue dari interface http.Request.
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		// * Konversi age dari string ke integer menggunakan strconv.Atoi.
		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
