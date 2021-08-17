package services

import (
	"github.com/kartik1998/bookstore-users-api/domain/users"
	"github.com/kartik1998/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return &user, err
	}
	return &user, nil
}
