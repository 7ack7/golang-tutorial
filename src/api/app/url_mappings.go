package app

import (
	"github.com/7ack7/golang-microservices/src/api/controllers/polo"
	"github.com/7ack7/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
