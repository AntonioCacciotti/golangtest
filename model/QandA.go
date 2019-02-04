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

//Iterate iterates over a slice
func (m *Questions) Iterate() []Questions {
	return questions
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
	Correct    bool   `json:"Correct"`
}

var questions []Questions

func getQuestion() []Questions {
	return questions
}

//GetAnswers get answer list
func GetAnswers(questionID int) Answers {
	log.Println("Loading questionID:", questionID, " size map:", len(answersMap))
	return answersMap[questionID]
}

//LoadQuestions from a file
func LoadQuestions() []Question {
	var questions Questions
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
