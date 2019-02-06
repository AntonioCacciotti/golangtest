package rest

import (
	"encoding/json"
	"fmt"
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
	router.PUT("/v1/:game/checkanswer", doCheckAnswer)
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
	for _, v := range answers {
		newAnswer := model.NewAnswer(v.ID, v.QuestionID, v.Text)
		resp = append(resp, newAnswer)
	}
	c.JSON(http.StatusOK, resp)
}

//CheckAnswer obj
type CheckAnswer struct {
	QID int    `json:"questionID"`
	AID int    `json:"answerID"`
	NID string `json:"nicknameID"`
}

//CheckAnswer call function to verify the user answer
func doCheckAnswer(c *gin.Context) {
	var checkAnswer CheckAnswer
	c.BindJSON(&checkAnswer)
	log.Println("json request body checkanswer:", checkAnswer)
	resp := model.VerifyUserAnswer(checkAnswer.QID, checkAnswer.AID, checkAnswer.NID)
	c.JSON(http.StatusOK, resp)
}

//ProcessAction returns one question and the related options TODO
func ProcessAction(c *gin.Context) {
	action := c.Query("action")
	nickname := c.Query("nickname")
	var result float64
	if action == "result" {
		result = model.GetScore(nickname)
		resp := fmt.Sprintf("%f", result)
		c.JSON(http.StatusOK, gin.H{"message": "Total correct answers:" + resp})
	}
	if action == "ranking" {
		log.Println("scores")
	}
}
