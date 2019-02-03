package model

//Question is a simple data struct
type Question struct {
	questionID int
	question   string
}

//LoadQuestion from a file
func LoadQuestion() map[int]string {
	questions := make(map[int]string)
	questions[1] = "ciao"
	return questions
}
