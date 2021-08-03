package controllers

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/users/api/domain/users"
	"github.com/users/api/services"
	"github.com/users/api/utils/errors"
)

func CreateUser(c *gin.Context) {
	//c.String(http.StatusNotImplemented, "")
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.RestErr{
			Messgae: "invalid json",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		log.Println(err)
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUsers(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
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
