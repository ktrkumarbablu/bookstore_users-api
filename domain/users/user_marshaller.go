package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id          int64  `json:"user_id"`
	DataCreated string `json:"date_created"`
	Status      string `json:"string"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DataCreated string `json:"date_created"`
	Status      string `json:"string"`
}

func (user *User) Mashall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DataCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateuser PrivateUser
	if err := json.Unmarshal(userJson, &privateuser); err != nil {
		return nil
	}
	return privateuser
}
