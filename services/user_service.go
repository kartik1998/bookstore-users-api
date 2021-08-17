package services

import "github.com/kartik1998/bookstore-users-api/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
