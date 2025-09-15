package dto

type FieldError struct {
	Field        string `json:"field"`
	ErrorMessage string `json:"error"`
}

type ValidationError struct {
	Errors FieldError `json:"errors"`
}

type RegistrationResponse struct {
	Status int    `json:"-"`
	Error  string `json:"error"`
}
