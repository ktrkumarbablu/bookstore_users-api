package users

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	userdb "github.com/users/api/datasources/mysql/user_db"
	utils "github.com/users/api/utils/crypt_utils"
	"github.com/users/api/utils/date_utils"
	"github.com/users/api/utils/errors"
)

const (
	QueryInsertUser = "INSERT INTO user(first_name,last_name,email,date_created,status,password) VALUES ( ?,?,?,?,?,?)"
	QueryGetUser    = "SELECT id,first_name,last_name,email,date_created FROM user where id=?"
	NoRowFound      = "no rows in result set"
	QueryUpdateUser = "UPDATE user SET first_name=?,last_name=?,email=? where id=?"
	QueryDeleteUser = "Delete from user where id=?"
	QueryByStatus   = "SELECT id,first_name,last_name,email,date_created,status,password FROM user where status=?"
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
	user.Password = utils.GetMD5(user.Password)
	insertResult, saveerr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, "active", user.Password)
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

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := userdb.ClientDB.Prepare(QueryByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	row, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer row.Close()
	result := make([]User, 0)
	for row.Next() {
		var user User
		if err := row.Scan(user.Id, user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NoRowFoundForData("no user are mathcing with this status")
	}
	return result, nil
}
