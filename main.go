package main

import (
	"database/sql"
	"fmt"
	"ilya-skoropad/user/config"
	"ilya-skoropad/user/internal/controller"
	"ilya-skoropad/user/internal/middleware"
	"ilya-skoropad/user/internal/repository"
	"log"
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

	wrappedMux := middleware.NewLoggerMiddleware(mux, log.Default())

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", conf.AppHost, conf.AppPort), wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
