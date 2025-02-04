package main

import (
	"sync"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Name string `json:"name"`
}

var (
	users      []User
	mu         sync.Mutex
	counter    int
	lastUpdate time.Time
)

func initializeStorage() {
	users = []User{}
	lastUpdate = time.Now()
}

func getAllUsers() []User {
	mu.Lock()
	defer mu.Unlock()
	return users
}

func addNewUser(newUser User) {
	mu.Lock()
	defer mu.Unlock()
	counter++
	newUser.ID = counter
	users = append(users, newUser)
	lastUpdate = time.Now()
}

func getLastUpdate() string {
	mu.Lock()
	defer mu.Unlock()
	return lastUpdate.Format(time.RFC3339)
}
