package controller

import (
	"encoding/json"
	"ilya-skoropad/user/internal/repository"
	"net/http"
)

type HealthResponse struct {
	Message string
}

type HealthController struct {
	repo repository.HealthRepository
}

func (h *HealthController) Handle(w http.ResponseWriter, r *http.Request) {
	err := h.repo.Ping()
	if err == nil {
		h.writeSuccess(w)
		return
	}

	h.writeError(w, err)
}

func (h *HealthController) writeSuccess(w http.ResponseWriter) {
	data, _ := json.Marshal(
		HealthResponse{
			Message: "ok",
		},
	)

	w.Write(data)
}

func (h *HealthController) writeError(w http.ResponseWriter, err error) {
	data, _ := json.Marshal(
		HealthResponse{
			Message: err.Error(),
		},
	)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(data)
}

func NewHealthController(repo repository.HealthRepository) *HealthController {
	return &HealthController{
		repo: repo,
	}
}
