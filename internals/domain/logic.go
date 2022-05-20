package domain

import "errors"

var (
	ErrItemNotFound   = errors.New("item not found")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInternalServer = errors.New("internal server error")
)

// User service
type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserRepository {
	return &userService{
		userRepo: userRepo,
	}
}

func (us userService) CreateUser(user User) (*User, error) {
	return us.userRepo.CreateUser(user)
}

func (us userService) GetUsers() (*[]User, error) {
	return us.userRepo.GetUsers()
}

func (us userService) GetUser(user_id string) (*User, error) {
	return us.userRepo.GetUser(user_id)
}

func (us userService) UpdateUser(user User) (*User, error) {
	return us.userRepo.UpdateUser(user)
}

func (us userService) DeleteUser(user_id string) error {
	return us.userRepo.DeleteUser(user_id)
}
