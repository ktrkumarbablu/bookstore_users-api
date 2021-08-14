package services

import (
	"github.com/users/api/domain/users"
	"github.com/users/api/utils/errors"
)

func CreateUsers(user users.User) (*users.User, *errors.RestErr) {
	//not imp
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: userId,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func Delete(user users.User) (*users.User, *errors.RestErr) {
	if err := user.DeleteUserByID(); err != nil {
		return nil, err
	}
	return nil, nil
}
