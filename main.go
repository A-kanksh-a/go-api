package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
	http.ListenAndServeTLS(":8080", "./cert/CA/localhost/localhost.crt", "./cert/CA/localhost/localhost.decrypted.key", r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello All")
}
