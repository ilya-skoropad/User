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

	h.userService.Register(request)
}

func NewRegistrationController(userService service.UserService) *registrationController {
	return &registrationController{
		userService: userService,
	}
}
