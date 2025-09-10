package dto

type RegistrationRequest struct {
	Nickname string `json:"nick"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}
