package service

import (
	"ilya-skoropad/user/internal/dto"
	"ilya-skoropad/user/internal/repository"
	"log"
)

type UserService interface {
	Register(request dto.RegistrationRequest)
}

type userService struct {
	repository repository.UserRepository
	logger     *log.Logger
}

func (s *userService) Register(request dto.RegistrationRequest) {
	guid, err := s.repository.FindGuidByLoginOrMail(request.Login, request.Email)
	if err != nil {
		panic(err)
	}

	s.logger.Println(guid)
}

func NewUserService(repository repository.UserRepository, logger *log.Logger) UserService {
	return &userService{
		repository: repository,
		logger:     logger,
	}
}
