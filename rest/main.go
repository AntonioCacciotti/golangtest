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
	router.POST("/v1/:game/userregistration", UserRegistration)
	router.Run(":8000")
}

var users = make(map[string]int)

//UserRegistration save user nickname inside a hashmap
func UserRegistration(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	users[user.Nickname] = 0
	log.Println(users)
	c.String(http.StatusOK, "il nickname e' %s", user.Nickname)
}

func UserRegistration(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	users[user.Nickname] = 0
	log.Println(users)
	c.String(http.StatusOK, "il nickname e' %s", user.Nickname)
}

//GetQuestions reads all questions from a file
func GetQuestions(c *gin.Context) {
	game := c.Param("game")
	action := c.Param("action")
	c.String(http.StatusOK, "ciao %s %s", game, action)
	log.Println(action)
	model.LoadQuestion()
}

//GetNextQuestions returns one question and the related options
func GetNextQuestions(c *gin.Context) {

}
