package main

import (
	"fmt"
	"ilya-skoropad/user/config"
	"ilya-skoropad/user/internal/controller"
	"log"
	"net/http"
)

func main() {
	conf := config.Get()

	mux := http.NewServeMux()

	healthController := controller.NewHealthController(
		log.Default(),
	)

	mux.HandleFunc("/health", healthController.Handle)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", conf.AppHost, conf.AppPort), mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
