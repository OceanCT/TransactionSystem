package main

import (
	"HomeworkForDB/Dao/DBView"
	"HomeworkForDB/Dao/Trigger"
	"HomeworkForDB/Router"
	"github.com/gin-gonic/gin"
)

func main() {
	Trigger.TriggerInit()
	DBView.ViewInit()
	r := gin.Default()
	Router.RegisterRouter(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(":13875")
}
