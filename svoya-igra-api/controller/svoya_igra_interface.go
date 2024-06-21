package controller

import "github.com/gin-gonic/gin"

type TopicControllerInterface interface {
	InsertTopic(*gin.Context)
	GetAllTopic(*gin.Context)
	GetOneTopic(*gin.Context)
	UpdateTopic(*gin.Context)
	DeleteTopic(*gin.Context)
}
