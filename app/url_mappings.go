package app

import "github.com/users/api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
