package users

import (
	"fmt"

	"github.com/kartik1998/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestErr(fmt.Sprintf("user with email %s already exists", user.Email))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("user %d already exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundReq(fmt.Sprintf("User id %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
