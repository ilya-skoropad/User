package controller

import (
	"encoding/json"
	"ilya-skoropad/user/internal/dto"
	"ilya-skoropad/user/internal/service"
	"io"
	"net/http"
)

type registrationController struct {
	userService service.UserService
}

func (h *registrationController) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var request dto.RegistrationRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		panic(err)
	}

	result := h.userService.Register(request)
	h.response(w, result)
}

func (h *registrationController) response(w http.ResponseWriter, resp dto.RegistrationResponse) {
	data, _ := json.Marshal(resp)

	w.WriteHeader(resp.Status)
	w.Write(data)
}

func NewRegistrationController(userService service.UserService) *registrationController {
	return &registrationController{
		userService: userService,
	}
}
