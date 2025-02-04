package main

import (
	"log"
)

func initializeReplication() {
	replicaUsers = []User{}
	lastUpdate = ""
	log.Println("Inicialización de REPLICACION completada")
}

func main() {
	initializeReplication()
	go shortPolling()
	go longPolling()
	log.Println("Servidor REPLICACION ejecutándose...")
	select {} // Mantiene el programa en ejecución
}
