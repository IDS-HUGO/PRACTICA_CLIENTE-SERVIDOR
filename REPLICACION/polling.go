package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var lastUpdate string

func shortPolling() {
	for {
		fetchUsers()
		time.Sleep(5 * time.Second) // Consulta cada 5 segundos
	}
}

func longPolling() {
	for {
		resp, err := http.Get("http://localhost:8080/lastModified")
		if err != nil {
			fmt.Println("Error verificando cambios:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		var data map[string]string
		json.Unmarshal(body, &data)

		if data["lastUpdate"] != lastUpdate {
			lastUpdate = data["lastUpdate"]
			fetchUsers()
		}
		time.Sleep(2 * time.Second) // Reduce la carga en el servidor
	}
}
