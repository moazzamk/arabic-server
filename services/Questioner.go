package services

import (
	"fmt"
	"math/rand"

	log "github.com/sirupsen/logrus"

	"arabic-server/models"
)

var verbs []models.Pronoun
var verbies []models.Verb

func init() {
	verbs = []models.Pronoun{}
	verbs = append(verbs, models.PresentPronouns...)
	verbs = append(verbs, models.PastPronouns...)

	verbies = []models.Verb{}
	verbies = append(verbies, models.PresentVerbs...)
	verbies = append(verbies, models.PastVerbs...)
}

type Questioner interface {
	GetQuestion() *models.Question
}

type questioner struct {}

func NewQuestioner() Questioner {
	return &questioner{}
}

func (r *questioner) GetQuestion() *models.Question {
	return r.identifyDoer();
/*
	if GetRandomNumber(2) < 1 {
		return r.identifyDoer()
	} else {
		return r.identifyFamily()
	}
 */
}

func (r *questioner) identifyDoer() *models.Question  {

	log.Info("Getting Doer Question")
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

	log.Info(ret, "Sending")
	return &ret
}

func (r *questioner) identifyFamily() *models.Question {
	ret := &models.Question{}

	log.Info("Getting Family Question")
	roots := generateRoot3()
	families := make([]models.Family, len(models.BigFamilies));
	copy(families, models.BigFamilies)


	// Get a family
	familyIndex := GetRandomNumber(len(models.BigFamilies))
	family := models.BigFamilies[familyIndex]

	// Get a tense and doer
	verbIndex := GetRandomNumber(len(verbies))
	verbie := verbies[verbIndex]

	middle := "";
	if verbie.Tense == "past" {
		middle = fmt.Sprintf(family.ActivePast, roots[0], roots[1], roots[2])
	} else {
		middle = fmt.Sprintf(family.ActivePresent, roots[0], roots[1], roots[2])
	}

	ret.QuestionText = `Identify the family: <span class=\"arabic\">` + verbie.Beginning + middle + verbie.Ending + `<span>`
	ret.CorrectAnswer = family.Name

	ret.Choices = generateFamilyChoices(ret.CorrectAnswer, familyIndex, families)
	log.Info(ret, "returning")

	return ret
}

func generateFamilyChoices(correctAnswer string, correctIndex int, ans []models.Family) []string {
	ret := []string{}

	numbers := make(map[int]struct{})
	numbers[correctIndex] = struct{}{}
	for ;len(numbers) < 4; {
		numbers[GetRandomNumber(len(ans))] = struct{}{}
	}

	for key, _ := range numbers {
		ret = append(ret, ans[key].Name)
	}

	return ret
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

	return roots
}

