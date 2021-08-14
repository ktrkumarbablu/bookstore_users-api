package controllers

import (
	"log"
	"strconv"

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
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		log.Println(err)

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
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be number")
		c.JSON(err.Status, err)
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, user)

	c.String(http.StatusNotImplemented, "")
}

func DeleteUser(c *gin.Context) {
	userid, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User does not exists for deletion")
		c.JSON(err.Status, err)
	}
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid josn body")
		c.JSON(restErr.Status, restErr)
	}
	user.Id = userid
	_, err := services.Delete(user)
	if err != nil {
		err := errors.NewBadRequestError("User id should be a number")
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, "user is deleted successfully")
}
func UpdateUserByOID(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User id should be a number")
		c.JSON(err.Status, err)
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid josn body")
		c.JSON(restErr.Status, restErr)
	}
	user.Id = userID
	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.FindUserByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)

}
