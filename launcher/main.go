package main

import (
	"golangtest/model"
	"golangtest/rest"
	"log"
)

func main() {
	log.Println("Loading questions and answers...")
	model.LoadQuestions()
	log.Println("Loading questions and answers...[OK]")
	log.Println("Strating server...")
	rest.StartServer()
	log.Println("Strating server...[OK]")

}
