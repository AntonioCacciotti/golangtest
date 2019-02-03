package model

//User is a simple data struct
type User struct {
	Nickname string `json:"nickname"`
	Score    int    `json:"score"`
}
