package app

import "github.com/users/api/controllers"

func mapUrls() {
	router.GET("v1/ping", controllers.Ping)
	router.POST("v1/createusers", controllers.CreateUser)
	router.GET("v1/user/:user_id", controllers.GetUserByID)
	router.GET("v1/users/search", controllers.FindUser)
}
