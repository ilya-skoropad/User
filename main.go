package main

import (
	"fmt"
	"ilya-skoropad/user/internal/controller"
	"log"
	"net/http"
)

func main() {
	healthController := controller.NewHealthController(
		log.Default(),
	)

	http.HandleFunc("/health", healthController.Handle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
