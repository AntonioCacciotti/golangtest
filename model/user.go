package model

import "log"

//User is a simple data struct
type User struct {
	Nickname string `json:"nickname"`
	Score    int    `json:"score"`
}

var users = make(map[string]int)

//AddUserNickname add a new user in a map
func AddUserNickname(nickname string) {
	users[nickname] = 0
	log.Println("Users: ", users)
}

//IncreamentUserScore func to incremente user poits
func IncreamentUserScore(nickname string) {
	score := users[nickname]
	score++
	users[nickname] = score
	log.Println("Users scores: ", users)
}
