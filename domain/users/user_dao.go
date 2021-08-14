package users

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	userdb "github.com/users/api/datasources/mysql/user_db"
	"github.com/users/api/utils/date_utils"
	"github.com/users/api/utils/errors"
)

const (
	QueryInsertUser = "INSERT INTO user(first_name,last_name,email,date_created) VALUES ( ?,?,?,?)"
	QueryGetUser    = "SELECT id,first_name,last_name,email,date_created FROM user where id=?"
	NoRowFound      = "no rows in result set"
	QueryUpdateUser = "UPDATE user SET first_name=?,last_name=?,email=? where id=?"
	QueryDeleteUser = "Delete from user where id=?"
)

var userDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	stmt, err := userdb.ClientDB.Prepare(QueryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), NoRowFound) {
			return errors.NoRowFoundForData("User does not exits")
		}
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := userdb.ClientDB.Prepare(QueryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveerr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveerr != nil {
		sqlerr, ok := saveerr.(*mysql.MySQLError)
		if ok != mysql.NewConfig().InterpolateParams {
			return errors.NewInternalServerError(err.Error())
		}
		fmt.Sprintf(sqlerr.Error())
		fmt.Sprintf(sqlerr.Message)

	}

	userid, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Sprintf("error when trying to save users: %s", userid)
		return errors.NewInternalServerError(err.Error())
	}
	user.Id = userid

	return nil
}

func (user User) Update() *errors.RestErr {
	stmt, err := userdb.ClientDB.Prepare(QueryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
func (user User) DeleteUserByID() *errors.RestErr {
	stmt, err := userdb.ClientDB.Prepare(QueryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
