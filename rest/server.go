package rest

import (
	"golangtest/model"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//StartServer starts the server
func StartServer() {
	router := gin.Default()
	router.GET("/v1/:game/questions", GetQuestions)
	router.GET("/v1/:game/answers", GetAnswers)
	router.GET("/v1/:game/end", ProcessAction)
	router.POST("/v1/:game/userregistration", UserRegistration)
	router.POST("/v1/:game/checkanswer", CheckAnswer)
	router.Run(":8000")
}

//UserRegistration save user nickname inside a hashmap
func UserRegistration(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	model.AddUserNickname(user.Nickname)
	c.String(http.StatusOK, "il nickname e' %s", user.Nickname)
}

//GetQuestions reads all questions from a file
func GetQuestions(c *gin.Context) {
	game := c.Param("game")
	if strings.ToLower(game) != "quiz" {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "wrong path!"})
	}
	action := c.Param("action")
	totalQ := len(model.GetQuestions().Questions)
	c.JSON(http.StatusOK, gin.H{"message": "try to answer to answer " + strconv.Itoa(totalQ) + " questions to get a beer!"})
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
	var nickname = c.Query("nickname")
	c.BindJSON(&answer)
	log.Println(answer)
	resp := model.VerifyUserAnswer(answer.QuestionID, answer.ID, nickname, answer.Correct)
	c.JSON(http.StatusOK, resp)
}

//ProcessAction returns one question and the related options
func ProcessAction(c *gin.Context) {
	action := c.Query("action")
	if action == "result" {
		log.Println("result")
	}
	if action == "scores" {
		log.Println("scores")
	}
}
