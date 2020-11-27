package models

type QuizState struct {
	CurrentQuestion *Question `json:"question"`
	Players []*Player `json:"players"`
	TotalQuestions int `json:"totalQuestions"`
}

func NewQuizState() *QuizState {
	return &QuizState{}
}
