package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Question is a simple data struct
type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
}

//Questions is a simple data struct
type Questions struct {
	Questions []Question `json:"questions"`
}

//LoadQuestion from a file
func LoadQuestion() map[int]string {
	var questions Questions
	jsonQuestions, err := os.Open("questions.json")
	byteValue, _ := ioutil.ReadAll(jsonQuestions)
	json.Unmarshal(byteValue, &questions)
	if err != nil {
		fmt.Println(err)
	}
	question2s := make(map[int]string)
	log.Println(questions)
	for i := 0; i < len(questions.Questions); i++ {
		fmt.Println("Question : " + questions.Questions[i].Question)
	}
	defer jsonQuestions.Close()
	return question2s
}
