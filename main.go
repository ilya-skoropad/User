package main

import (
	"database/sql"
	"fmt"
	"ilya-skoropad/user/config"
	"ilya-skoropad/user/internal/controller"
	"ilya-skoropad/user/internal/repository"
	"net/http"
)

func main() {
	conf := config.Get()

	mux := http.NewServeMux()

	connection, err := sql.Open("postgres", conf.DbCon)
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	healthController := controller.NewHealthController(
		repository.NewHealthRepository(connection),
	)

	mux.HandleFunc("/health", healthController.Handle)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", conf.AppHost, conf.AppPort), mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
