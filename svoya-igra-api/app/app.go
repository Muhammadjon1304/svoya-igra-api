package app

import (
	"database/sql"
	"github.com/muhammadjon1304/svoya-igra-api/controller"
	"github.com/muhammadjon1304/svoya-igra-api/db"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := db.Connectdb()
	a.DB = db
}
func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewTopicController(a.DB)
	r.POST("/topic", controller.InsertTopic)
	r.GET("/topic", controller.GetAllTopic)
	r.GET("/topic/:id", controller.GetOneTopic)
	r.PUT("/topic/:id", controller.UpdateTopic)
	r.DELETE("/topic/:id", controller.DeleteTopic)
	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
