package model

//User is a simple data struct
type User struct {
	FirstName string
	LastName  string
	Score     int
	Rank      int
}

//NewUser create new user
func NewUser(name string, surname string, score int, rank int) User {
	return User{name, surname, score, rank}
}
