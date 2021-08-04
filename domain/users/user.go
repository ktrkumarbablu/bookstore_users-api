package users

import (
	"strings"

	"github.com/users/api/utils/errors"
)

type User struct {
	Id          int64
	FirstName   string `json:"id"`
	LastName    string `json:"firstName"`
	Email       string `json:"email"`
	DateCreated string `json:"datecreated"`
}

func Validate(users *User) *errors.RestErr {
	users.Email = strings.TrimSpace(strings.ToLower(users.Email))
	if users.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
