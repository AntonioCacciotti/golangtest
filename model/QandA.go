package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Question is a simple obj
type Question struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

//Questions array
type Questions struct {
	Questions []Question `json:"questions"`
}

//Answers array
type Answers struct {
	Answers []Answer `json:"answers"`
}

//Answer is a simple obj
type Answer struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"questionID"`
	Text       string `json:"text"`
	Correct    bool   `json:"correct"`
}

//GetAnswers get answer list
func GetAnswers(questionID int) Answers {
	log.Println("Loading questionID:", questionID, " size map:", len(answersMap))
	return answersMap[questionID]
}

//NewAnswer builds new obj for response
func NewAnswer(ID int, questionID int, text string) Answer {
	return Answer{ID, questionID, text, false}
}

//VerifyUserAnswer builds new obj for response
func VerifyUserAnswer(questionID int, answerID int, username string) bool {
	log.Println("verify answer...")
	answer := answersMap[questionID]
	for _, v := range answer.Answers {
		if v.ID == answerID {
			if v.Correct == true {
				IncreamentUserScore(username)
				return true
			}
		}
	}
	return false
}

var questions Questions

//GetQuestions return all the questions present in the questions.json file
func GetQuestions() Questions {
	return questions
}

//GetNextQuestionsWithAnswers return all the questions present in the questions.json file
func GetNextQuestionsWithAnswers(questionID int) (Question, Answers) {
	var question Question
	var answers Answers
	for _, v := range questions.Questions {
		if v.ID == questionID {
			return v, GetAnswers(v.ID)
		}
	}
	return question, answers
}

//LoadQuestions from a file
func LoadQuestions() []Question {

	jsonQuestions, err := os.Open("questions.json")
	byteValue, _ := ioutil.ReadAll(jsonQuestions)
	json.Unmarshal(byteValue, &questions)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(questions)
	for i, k := range questions.Questions {
		log.Println("Question : ", k.ID, "Loaded!", i)
		LoadAnswers(k.ID)
	}
	defer jsonQuestions.Close()
	return questions.Questions
}

var answersMap = make(map[int]Answers)

//LoadAnswers from a file
func LoadAnswers(questionID int) map[int]Answers {
	var answers Answers
	jsonAnswers, err := os.Open("answers.json")
	byteValue, _ := ioutil.ReadAll(jsonAnswers)
	json.Unmarshal(byteValue, &answers)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(answers)
	for _, v := range answers.Answers {
		if v.ID == questionID {
			log.Println("QuestionId:", questionID, "Answer : ", v.Text, "Loaded!")
			answersMap[questionID] = answers
			log.Println("Answer added in map : ", answersMap)
		}
	}
	defer jsonAnswers.Close()
	return answersMap
}
