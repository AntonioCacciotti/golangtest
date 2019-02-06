package model

import (
	"fmt"
	"log"
	"math"
	"sort"
)

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

//GetScore get user score and return %
func GetScore(nickname string) float64 {
	score := float64(users[nickname])
	var result, q float64
	q = float64(len(GetQuestions().Questions))
	result = (score / q) * 100
	fmt.Println("score for nickname", nickname, " is:", toFixed(result, 0))
	return toFixed(result, 0)
}

//GetRanking get ranking of users
func GetRanking() map[int][]string {
	n := map[int][]string{}
	var a []int
	for k, v := range users {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	}
	return n
}
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
