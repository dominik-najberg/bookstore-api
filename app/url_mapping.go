package app

import (
	"github.com/dominik-najberg/bookstore-users-api/controller"
)

func mapUrls() {
	router.GET("/ping", controller.Ping)
}