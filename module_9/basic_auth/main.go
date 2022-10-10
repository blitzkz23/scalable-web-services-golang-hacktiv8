package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ! Basic auth

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server is running on port 9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	// Cek apakah request basic auth
	if !Auth(w, r) {
		return
	}
	// Cek hanya boleh method get
	if !AllowOnlyGET(w, r) {
		return
	}

	// Cek apakah url punya parameter id
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
