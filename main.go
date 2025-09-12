package main

import (
	"database/sql"
	"fmt"
	"ilya-skoropad/user/config"
	"ilya-skoropad/user/internal/controller"
	"ilya-skoropad/user/internal/middleware"
	"ilya-skoropad/user/internal/repository"
	"ilya-skoropad/user/internal/service"
	"log"
	"net/http"
)

func main() {
	conf := config.Get()

	connection, err := sql.Open("postgres", conf.DbCon)
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	logger := log.Default()
	healthController := controller.NewHealthController(
		repository.NewHealthRepository(connection),
	)

	registrationController := controller.NewRegistrationController(
		service.NewUserService(
			repository.NewUserRepository(connection),
			logger,
		),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthController.Handle)
	mux.HandleFunc("POST /api/register", registrationController.Handle)

	wrappedMux := middleware.NewLoggerMiddleware(mux, logger)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", conf.AppHost, conf.AppPort), wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
