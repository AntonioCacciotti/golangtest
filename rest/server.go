package rest

import (
	"golangtest/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//StartServer starts the server
func StartServer() {
	router := gin.Default()
	router.GET("/v1/:game/questions", GetQuestions)
	//router.GET("/v1/:game/:action", GetQuestions)
	router.GET("/v1/:game/answers", GetAnswers)
	router.POST("/v1/:game/userregistration", UserRegistration)
	router.POST("/v1/:game/checkanswer", CheckAnswer)
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

//GetQuestions reads all questions from a file
func GetQuestions(c *gin.Context) {
	game := c.Param("game")
	action := c.Param("action")
	c.String(http.StatusOK, "ciao %s %s", game, action)
	log.Println(action)
}

//GetNextQuestions returns one question and the related options
func GetNextQuestions(c *gin.Context) {
	nextQuestion := c.Query("next")
	log.Println(nextQuestion)
}

//GetAnswers returns all the question' answers
func GetAnswers(c *gin.Context) {
	questionID := c.Query("questionID")
	log.Println("questionId: ", questionID)
	qID, err := strconv.Atoi(questionID)
	if err != nil {
		log.Fatalln("convertsion string to int FATAL ERROR")
	}
	answers := model.GetAnswers(qID)
	resp := []model.Answer{}
	for _, v := range answers.Answers {
		newAnswer := model.NewAnswer(v.ID, v.QuestionID, v.Text)
		resp = append(resp, newAnswer)
	}
	c.JSON(http.StatusOK, resp)
}

//CheckAnswer call function to verify the user answer
func CheckAnswer(c *gin.Context) {
	var answer model.Answer
	c.BindJSON(&answer)
	log.Println(answer)
	resp := model.VerifyUserAnswer(answer.QuestionID, answer.ID, answer.Correct)
	c.JSON(http.StatusOK, resp)
}
