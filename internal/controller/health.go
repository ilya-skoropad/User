package controller

import (
	"encoding/json"
	"ilya-skoropad/user/internal/repository"
	"net/http"
)

type HealthResponse struct {
	Message string
}

type healthController struct {
	repo repository.HealthRepository
}

func (h *healthController) Handle(w http.ResponseWriter, r *http.Request) {
	err := h.repo.Ping()
	if err == nil {
		h.writeSuccess(w)
		return
	}

	h.writeError(w, err)
}

func (h *healthController) writeSuccess(w http.ResponseWriter) {
	data, _ := json.Marshal(
		HealthResponse{
			Message: "ok",
		},
	)

	w.Write(data)
}

func (h *healthController) writeError(w http.ResponseWriter, err error) {
	data, _ := json.Marshal(
		HealthResponse{
			Message: err.Error(),
		},
	)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(data)
}

func NewHealthController(repo repository.HealthRepository) *healthController {
	return &healthController{
		repo: repo,
	}
}
