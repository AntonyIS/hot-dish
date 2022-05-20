package domain

type UserRepository interface {
	CreateUser(user User) (*User, error)
	GetUser(user_id string) (*User, error)
	GetUsers() (*[]User, error)
	UpdateUser(user User) (*User, error)
	DeleteUser(user_id string) error
}
