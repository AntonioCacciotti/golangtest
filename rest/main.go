package main

import (
	"golangtest/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/v1/:game/:action", GetQuestions)
	router.POST("v1/:game/registeruser")
	router.Run(":8000")
}

//GetQuestions reads all questions from a file
func GetQuestions(c *gin.Context) {
	game := c.Param("game")
	action := c.Param("action")
	c.String(http.StatusOK, "ciao %s %s", game, action)
	log.Println(action)
	model.LoadQuestion()
}
