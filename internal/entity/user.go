package entity

import (
	"ilya-skoropad/user/internal/enum"
	"time"
)

type Role struct {
	Id   int
	Name string
}

type User struct {
	Guid          string
	State         enum.State
	Role          Role
	Nickname      string
	Login         string
	Email         string
	Password      string
	CreatedAt     time.Time
	ActivationKey string
	ActivatedAt   time.Time
	LastLogineAt  time.Time
}
