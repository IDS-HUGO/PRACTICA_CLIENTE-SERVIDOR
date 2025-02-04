package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAllUsers())
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}
	addNewUser(newUser)
	w.WriteHeader(http.StatusCreated)
}

func lastModified(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"lastUpdate": getLastUpdate()})
}
