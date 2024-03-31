package repository

import (
	"gorm.io/gorm"
	"linktree/internal/service"
)

type UserEntity struct {
	gorm.Model
	Email          string
	HashedPassword string
	Username       string
	PhoneNumber    string
	Salt           string
}

func (u UserEntity) TableName() string {
	return "users"
}

func toSvcUserEntity(entity UserEntity) service.UserEntity {
	return service.UserEntity{
		ID:          entity.ID,
		Email:       entity.Email,
		Username:    entity.Username,
		PhoneNumber: entity.PhoneNumber,
	}
}

func fromSvcCreateUserRequest(request service.CreateUserRequest) UserEntity {
	return UserEntity{
		Email:          request.Email,
		HashedPassword: request.HashedPassword,
		Username:       request.Username,
		PhoneNumber:    request.PhoneNumber,
		Salt:           request.Salt,
	}
}

func toSvcCreateUserResponse(entity UserEntity) service.CreateUserResponse {
	return service.CreateUserResponse{
		UserEntity: toSvcUserEntity(entity),
	}
}

func toSvcGetUserResponse(entity UserEntity) service.GetUserResponse {
	return service.GetUserResponse{
		UserEntity:     toSvcUserEntity(entity),
		HashedPassword: entity.HashedPassword,
		Salt:           entity.Salt,
	}
}

func toSvcGetAllUsersResponse(entities []UserEntity) service.GetAllUsersResponse {
	svcUserEntities := make([]service.UserEntity, 0, len(entities))

	for _, entity := range entities {
		svcUserEntities = append(svcUserEntities, toSvcUserEntity(entity))
	}

	return service.GetAllUsersResponse{UserEntities: svcUserEntities}
}
