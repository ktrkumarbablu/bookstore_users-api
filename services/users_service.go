package services

import (
	"github.com/users/api/domain/users"
	"github.com/users/api/utils/errors"
)

func CreateUsers(user users.User) (*users.User, *errors.RestErr) {
	//not imp
	return &user, nil
}
