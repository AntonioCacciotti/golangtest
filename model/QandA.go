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
func GetAnswers(questionID int) []Answer {
	log.Println("Loading questionID:", questionID, " size map:", len(answersMap))
	return answersMap[questionID]
}

//NewAnswer builds new obj for response
func NewAnswer(ID int, questionID int, text string) Answer {
	return Answer{ID, questionID, text, false}
}

//VerifyUserAnswer builds new obj for response
func VerifyUserAnswer(questionID int, answerID int, username string) bool {
	log.Println("verify answer... questionId", questionID)
	log.Println("verify answersMap:", answersMap)
	answer := answersMap[questionID]
	log.Println("verify answer:", answer)
	for _, v := range answer {
		log.Println("verify answerID:", v.ID, "answer input:", answerID)
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
func GetNextQuestionsWithAnswers(questionID int) (Question, []Answer) {
	var question Question
	var answers []Answer
	for _, v := range questions.Questions {
		if v.ID == questionID {
			fmt.Println("id from qiestions :", v.ID, "id from request:", questionID)
			fmt.Println("answers from:", v.ID, "-", GetAnswers(v.ID))
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

var answersMap = make(map[int][]Answer)

//LoadAnswers from a file
func LoadAnswers(questionID int) {
	var answers Answers
	jsonAnswers, err := os.Open("answers.json")
	byteValue, _ := ioutil.ReadAll(jsonAnswers)
	json.Unmarshal(byteValue, &answers)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(answers)
	for _, v := range answers.Answers {
		if v.QuestionID == questionID {
			//the bug is on answersMap I have to use pointers
			log.Println("checking bug QuestionId:", questionID, "v.ID", v.Text, "Loaded!")
			log.Println("QuestionId:", questionID, "Answer : ", v.Text, "Loaded!")
			answersMap[questionID] = append(answersMap[questionID], v)
			log.Println("Answer added in map : ", answersMap)
		}
	}
	defer jsonAnswers.Close()

}
