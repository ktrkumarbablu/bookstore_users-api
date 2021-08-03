package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/users/api/domain/users"
	"github.com/users/api/services"
)

func CreateUser(c *gin.Context) {
	//c.String(http.StatusNotImplemented, "")
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(bytes))
	if err != nil {
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		return
	}
	result, saveErr := services.CreateUsers(user)
	if saveErr != nil {
		return
	}

	c.JSON(http.StatusCreated, result)

}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "")
}
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "")
}
func GetUserByID(c *gin.Context) {
	c.String(http.StatusNotImplemented, "")
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "")
}
