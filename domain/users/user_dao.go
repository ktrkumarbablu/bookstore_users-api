package users

import (
	"fmt"
	"time"

	userdb "github.com/users/api/datasources/mysql/user_db"
	"github.com/users/api/utils/errors"
)

var userDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	if err := userdb.ClientDB.Ping(); err != nil {
		panic(err)
	}
	result := userDB[user.Id]
	if result != nil {
		return errors.NotFoundError("user doesn't exist")
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email id is already exists"))
		}
		return errors.NewBadRequestError(fmt.Sprintf("email is already exist in the database"))
	}
	now := time.Now()
	userDB[user.Id] = user
	user.DateCreated = now.String()
	return nil
}

func (user User) Update() *errors.RestErr {
	return nil
}
