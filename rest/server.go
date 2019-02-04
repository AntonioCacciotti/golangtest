package rest

import (
	"encoding/json"
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
	router.PUT("/v1/:game/checkanswer", CheckAnswer)
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
	nextQuestion := c.Query("next")
	if strings.ToLower(game) != "quiz" {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "wrong path!"})
	}
	if nextQuestion == "" {
		totalQ := len(model.GetQuestions().Questions)
		c.JSON(http.StatusOK, gin.H{"message": "try to answer to answer " + strconv.Itoa(totalQ) + " questions to get a beer!"})
	} else {
		GetNextQuestions(c, nextQuestion)
	}
}

//GetNextQuestions returns one question and the related options
func GetNextQuestions(c *gin.Context, next string) {
	nextQuestion, err := strconv.Atoi(next)
	if err != nil {
		log.Fatalln("convertsion string to int FATAL ERROR")
	}
	question, answers := model.GetNextQuestionsWithAnswers(nextQuestion)
	qToJSON, err := json.Marshal(question)
	aToJSON, err := json.Marshal(answers)
	q := string(qToJSON)
	a := string(aToJSON)
	c.JSON(http.StatusOK, gin.H{"question": q, "answers": a})

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

type checkAnswer struct {
	qID int    `json:"questionId" binding:"required"`
	aID int    `json:"answerId" binding:"required"`
	nID string `json:"nicknameID" binding:"required"`
}

//CheckAnswer call function to verify the user answer
func CheckAnswer(c *gin.Context) {
	var checkAnswer checkAnswer
	c.BindJSON(&checkAnswer)
	log.Println(checkAnswer)
	resp := model.VerifyUserAnswer(checkAnswer.qID, checkAnswer.aID, checkAnswer.nID)
	c.JSON(http.StatusOK, resp)
}

//ProcessAction returns one question and the related options TODO
func ProcessAction(c *gin.Context) {
	action := c.Query("action")
	if action == "result" {
		log.Println("result")
	}
	if action == "scores" {
		log.Println("scores")
	}
}
