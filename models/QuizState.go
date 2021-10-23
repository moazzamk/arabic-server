package models

// QuizState is the current state of a quiz. It's function is to let clients know that
// something has changed
type QuizState struct {
	CurrentQuestion *Question `json:"question"`
	Players []*Player `json:"players"`
	TotalQuestions int `json:"totalQuestions"`
}

func NewQuizState() *QuizState {
	return &QuizState{}
}
