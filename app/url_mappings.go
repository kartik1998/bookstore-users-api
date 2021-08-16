package app

import "github.com/kartik1998/bookstore-users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
