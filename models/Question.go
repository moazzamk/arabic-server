package models

type Question struct {
	QuestionText string `json:"question-text"`
	Choices []string `json:"choices"`
	CorrectAnswer string `json:"-"`
}


