package controller

import "github.com/sirupsen/logrus"

type UserService interface {
}

type User struct {
	logger  *logrus.Entry
	userSvc UserService
}

func NewUser(logger *logrus.Entry, userSvc UserService) *User {
	return &User{
		logger:  logger,
		userSvc: userSvc,
	}
}
