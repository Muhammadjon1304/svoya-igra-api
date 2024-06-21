package repository

import "github.com/muhammadjon1304/svoya-igra-api/model"

type TopicRepositoryInterface interface {
	InsertTopic(model.PostTopic) bool
	GetAllTopic() []model.Topic
	GetOneTopic(uint) model.Topic
	UpdateTopic(uint, model.PostTopic) model.Topic
	DeleteTopic(uint2 uint) bool
}
