package main

import (
	"backend-majoo-test/connection"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var pg connection.PgDatabase

func init() {
	pg = connection.PgDatabase{}
}

func main() {
	destination := "host=localhost user=postgres password=qweasd123 dbname=majoo port=5449 sslmode=disable"
	db, err := pg.NewInstance(destination)
	if err != nil {
		panic(err)
	}

	log.Printf("Database is running: %v", db.ConnPool)

	connection.Seed(db)

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
