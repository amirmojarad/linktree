package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"linktree/internal/service"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (r User) CreateUser(ctx context.Context, req service.CreateUserRequest) (service.CreateUserResponse, error) {
	userEntity := fromSvcCreateUserRequest(req)

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&userEntity).Error
	}); err != nil {
		return service.CreateUserResponse{}, errors.WithStack(err)
	}

	return toSvcCreateUserResponse(userEntity), nil
}

func (r User) FirstOrCreate(ctx context.Context, req service.CreateUserRequest) (service.CreateUserResponse, error) {
	userEntity := fromSvcCreateUserRequest(req)

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("username = ?", req.Username).Where("email = ?", req.Email).FirstOrCreate(&userEntity).Error
	}); err != nil {
		return service.CreateUserResponse{}, err
	}

	return toSvcCreateUserResponse(userEntity), nil
}

func (r User) GeUser(ctx context.Context, f service.GetUserFilter) (service.GetUserResponse, error) {
	var userEntity UserEntity

	query := r.db.WithContext(ctx).Model(&UserEntity{})
	query = appendConditionsToQuery(query, f.ConvertToMap())

	if err := query.Find(&userEntity).Error; err != nil {
		return service.GetUserResponse{}, errors.WithStack(err)
	}

	return toSvcGetUserResponse(userEntity), nil
}

func (r User) GetAllUsers(ctx context.Context, f service.GetAllUsersFilter) (service.GetAllUsersResponse, error) {
	entities := make([]UserEntity, 0)

	query := r.db.WithContext(ctx).Model(&UserEntity{})
	query = appendConditionsToQuery(query, f.ConvertToMap())

	if err := query.Find(entities).Error; err != nil {
		return service.GetAllUsersResponse{}, errors.WithStack(err)
	}

	return toSvcGetAllUsersResponse(entities), nil
}
