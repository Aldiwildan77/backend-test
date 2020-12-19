package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// setup router
	router := mux.NewRouter()
	router.HandleFunc("/", index)

	// run server
	if err := http.ListenAndServe(":8009", router); err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}
