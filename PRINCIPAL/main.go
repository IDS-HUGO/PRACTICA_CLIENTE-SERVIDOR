package main

import (
	"log"
	"net/http"
)

func main() {
	initializeStorage()
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/add", addUser)
	http.HandleFunc("/lastModified", lastModified)

	log.Println("Servidor PRINCIPAL ejecutándose en :8080")
	http.ListenAndServe(":8080", nil)
}
