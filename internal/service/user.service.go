package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"linktree/config"
)

type CreateUpdateUserRepository interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
	FirstOrCreate(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
}

type RetrieveUserRepository interface {
	GeUser(ctx context.Context, f GetUserFilter) (GetUserResponse, error)
	GetAllUsers(ctx context.Context, f GetAllUsersFilter) (GetAllUsersResponse, error)
}

type User struct {
	cfg      *config.AppConfig
	logger   *logrus.Entry
	tracer   trace.Tracer
	userRepo CreateUpdateUserRepository
}

func NewUser(
	cfg *config.AppConfig,
	logger *logrus.Entry,
	tracer trace.Tracer,
	userRepo CreateUpdateUserRepository,
) *User {
	return &User{
		cfg:      cfg,
		logger:   logger,
		tracer:   tracer,
		userRepo: userRepo,
	}
}
