package service

type UserEntity struct {
	ID          uint
	Email       string
	Username    string
	PhoneNumber string
}

type CreateUserResponse struct {
	UserEntity
}

type CreateUserRequest struct {
	Username       string
	Email          string
	PhoneNumber    string
	HashedPassword string
	Salt           string
}

type GetAllUsersResponse struct {
	UserEntities []UserEntity
}

type GetAllUsersFilter struct {
	Emails       []string
	PhoneNumbers []string
	Usernames    []string
}

type GetUserResponse struct {
	UserEntity
	HashedPassword string
	Salt           string
}

type GetUserFilter struct {
	Username    *string
	Email       *string
	PhoneNumber *string
}

func (f GetUserFilter) ConvertToMap() map[string]any {
	return map[string]any{
		"username = ?":     f.Username,
		"email = ?":        f.Email,
		"phone_number = ?": f.PhoneNumber,
	}
}

func (f GetAllUsersFilter) ConvertToMap() map[string]any {
	return map[string]any{
		"username in ?":     f.Usernames,
		"email in ?":        f.Emails,
		"phone_number in ?": f.PhoneNumbers,
	}
}
