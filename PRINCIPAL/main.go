package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	initializeStorage()

	// Rutas para CRUD
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", getUser).Methods("GET")
	router.HandleFunc("/add", addUser).Methods("POST")
	router.HandleFunc("/lastModified", lastModifiedTime).Methods("GET")

	// Rutas para polling
	router.HandleFunc("/polling/short", shortPolling).Methods("GET")
	router.HandleFunc("/polling/long", longPolling).Methods("GET")

	log.Println("Servidor PRINCIPAL ejecutándose en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
