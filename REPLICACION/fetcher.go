package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Name string `json:"name"`
}

var replicaUsers []User

func fetchUsers() {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Println("Error al obtener usuarios:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var users []User
	json.Unmarshal(body, &users)

	replicaUsers = users
	fmt.Println("Usuarios replicados:", replicaUsers)
}
