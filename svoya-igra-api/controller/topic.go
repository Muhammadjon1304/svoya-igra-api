package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/svoya-igra-api/model"
	"github.com/muhammadjon1304/svoya-igra-api/repository"
)

type TopicController struct {
	Db *sql.DB
}

func NewTopicController(db *sql.DB) TopicControllerInterface {
	return &TopicController{Db: db}
}

func (t *TopicController) DeleteTopic(c *gin.Context) {
	DB := t.Db
	var uri model.TopicUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTopicRepository(DB)
	delete := repository.DeleteTopic(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "deleted topic succesfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "deleted topic failed"})
		return
	}
}

func (t *TopicController) GetAllTopic(c *gin.Context) {
	DB := t.Db
	repository := repository.NewTopicRepository(DB)
	get := repository.GetAllTopic()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "msg": "got topic succesfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "got topic failed"})
		return
	}
}

func (t *TopicController) GetOneTopic(c *gin.Context) {
	DB := t.Db
	var uri model.TopicUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTopicRepository(DB)
	get := repository.GetOneTopic(uri.ID)
	if (get != model.Topic{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get topic successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "topic not found"})
		return
	}

}

func (t *TopicController) InsertTopic(c *gin.Context) {
	DB := t.Db
	var post model.PostTopic
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTopicRepository(DB)
	insert := repository.InsertTopic(post)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert topic successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert topic failed"})
		return
	}

}

func (t *TopicController) UpdateTopic(c *gin.Context) {
	DB := t.Db
	var post model.PostTopic
	var uri model.TopicUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTopicRepository(DB)
	update := repository.UpdateTopic(uri.ID, post)
	if (update != model.Topic{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "msg": "update topic successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "msg": "update topic failed"})
		return
	}
}
