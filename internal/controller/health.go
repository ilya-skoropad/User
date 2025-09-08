package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string
}

type HealthController struct {
	logger *log.Logger
}

func (h *HealthController) Handle(w http.ResponseWriter, r *http.Request) {
	response := &HealthResponse{
		Status: "OK",
	}

	data, err := json.Marshal(response)
	if err != nil {
		h.logger.Fatal(err.Error())
	}

	w.Write(data)
}

func NewHealthController(logger *log.Logger) *HealthController {
	return &HealthController{
		logger: logger,
	}
}
