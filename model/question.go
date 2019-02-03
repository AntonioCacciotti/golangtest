package model

import (
	"fmt"
	"log"
	"os"
)

//Question is a simple data struct
type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
}

//LoadQuestion from a file
func LoadQuestion() map[int]string {
	questions, err := os.Open("../resources/questions.json")
	if err != nil {
		fmt.Println(err)
	}
	question2s := make(map[int]string)
	log.Println(questions)
	defer questions.Close()
	return question2s
}
