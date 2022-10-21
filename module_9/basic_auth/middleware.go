package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "gotham"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	// Cek apakah request basi auth, jika ok maka lanjutkan otentikasi username dan password
	username, password, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`invalid username or password`))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`method not allowed`))
		return false
	}

	return true
}
