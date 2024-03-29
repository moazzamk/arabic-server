package services

import (
	"arabic-server/models"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func Test_it_generates_3_letter_roots(t *testing.T)  {
	Equal(t, len(generateRoot3()), 3)
}

func Test_it_gets_family_question(t *testing.T) {
	obj := &questioner{}
	rs := obj.identifyFamily()
	Equal(t, len(rs.Choices), 4)
	NotEmpty(t, rs.QuestionText)
	NotEmpty(t, rs.CorrectAnswer)
}

func Test_it_generate_4_choices_with_right_answer(t *testing.T) {
	found := false
	rs := generateAnswers("hello", models.PastPronouns)
	for i := 0; i < len(rs); i++ {
		if rs[i] == "hello" {
			found = true
		}
	}

	Equal(t, true, found)
	Equal(t, 4, len(rs))
}

func Test_it_gets_a_question_with_4_choices(t *testing.T) {
	subject := NewQuestioner()
	q := subject.GetQuestion()

	NotEmpty(t, q.QuestionText)
	NotEmpty(t, q.CorrectAnswer)
	Equal(t, 4, len(q.Choices))

	found := false
	for i := 0; i < len(q.Choices); i++ {
		if q.Choices[i] == q.CorrectAnswer {
			found = true
		}
	}

	Equal(t, true, found)
}
