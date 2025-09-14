package service

import (
	"ilya-skoropad/user/internal/dto"
	"ilya-skoropad/user/internal/entity"
	"ilya-skoropad/user/internal/repository"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(request dto.RegistrationRequest) dto.RegistrationResponse
}

type userService struct {
	repository repository.UserRepository
	logger     *log.Logger
}

func (s *userService) Register(request dto.RegistrationRequest) dto.RegistrationResponse {
	err := s.repository.CheckExistsByLoginOrMail(request.Login, request.Email)
	if err == nil {
		return dto.RegistrationResponse{
			Status: http.StatusConflict,
			Error:  "Login or password is takken",
		}
	}

	if err.Error() != "sql: no rows in result set" {
		panic(err)
	}

	pass, err := s.hashPassword(request.Password)
	if err != nil {
		panic(err)
	}

	user := entity.User{
		Login:    request.Login,
		Nickname: request.Nickname,
		Email:    request.Email,
		Password: pass,
	}

	err = s.repository.Save(user)
	if err != nil {
		panic(err)
	}

	return dto.RegistrationResponse{
		Status: http.StatusOK,
	}
}

func (s *userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *userService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUserService(repository repository.UserRepository, logger *log.Logger) UserService {
	return &userService{
		repository: repository,
		logger:     logger,
	}
}
