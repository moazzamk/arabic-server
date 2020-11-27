package services

import (
	"math/rand"
	"fmt"

	log "github.com/sirupsen/logrus"

	"arabic-server/models"
)
var verbs []models.Pronoun

func init() {
	verbs = []models.Pronoun{}
	verbs = append(verbs, models.PresentPronouns...)
	verbs = append(verbs, models.PastPronouns...)
}

type Questioner interface {
	GetQuestion() *models.Question
}

type questioner struct {

}

func NewQuestioner() Questioner {
	return &questioner{}
}

func (r *questioner) GetQuestion() *models.Question {
	log.Info("Getting Question")
	fi3l := make([]models.Pronoun, len(verbs));
	copy(fi3l, verbs)

	ret := models.Question{}
	roots := generateRoot3()
	index := GetRandomNumber(len(models.ArabicLetters) - 1)
	item := fi3l[index];
	fi3l = append(fi3l[0:index], fi3l[index+1:]...)

	question := fmt.Sprintf("Identify the inside doer: " + item.Pattern, roots[0], roots[1], roots[2])

	ret.QuestionText = question
	ret.CorrectAnswer = fmt.Sprintf("%s (%s)", item.Pronoun, item.Tense)
	ret.Choices = generateAnswers(ret.CorrectAnswer, fi3l)

	return &ret
}

func generateAnswers(correctAnswer string, ans []models.Pronoun) []string {
	ret := []string{
		correctAnswer,
	}

	for i := 0; i < 3; i++ {
		index := GetRandomNumber(len(ans) - 1)
		item := ans[index]
		ret = append(ret, fmt.Sprintf("%s (%s)", item.Pronoun, item.Tense))

		ans = append(ans[0:index], ans[index+1:]...)
	}
	rand.Shuffle(len(ret), func(i, j int) { ret[i], ret[j] = ret[j], ret[i]})

	return ret
}

func generateRoot3() []string {
	roots := []string{}

	roots = append(roots, models.ArabicLetters[GetRandomNumber(len(models.ArabicLetters) - 1)])
	roots = append(roots, models.ArabicLetters[GetRandomNumber(len(models.ArabicLetters) - 1)])
	roots = append(roots, models.ArabicLetters[GetRandomNumber(len(models.ArabicLetters) - 1)])

	return roots;
}

