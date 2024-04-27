package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}
		paramName := r.URL.Query().Get("name")
		if paramName == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
			return
		}

		exist := IsNameExists(paramName)
		if !exist {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Name is exists"))

	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
