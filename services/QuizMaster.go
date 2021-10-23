package services

import (
	"arabic-server/models"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strings"
)

var quizzes = map[string]*models.Quiz{}
var clients = map[*websocket.Conn]string{} // quiz ID is the value


// QuizMaster is responsible for creating quizzes, maintaining their state
type QuizMaster interface {
	QuizExists(id string) bool
	CreateQuiz(id string, optionValues url.Values)

	CheckAnswer(id string, connection *websocket.Conn, message *models.Message)

	AddPlayer(id string, connection *websocket.Conn, player *models.Player)
	DeletePlayer(id string, connection *websocket.Conn)
}

// NewQuizMaster returns new instance of quizMaster
func NewQuizMaster(questioner Questioner) QuizMaster {
	return &quizMaster{
		questioner: questioner,
	}
}

type quizMaster struct {
	questioner Questioner
}

func (r *quizMaster) CreateQuiz(id string, optionValues url.Values) {
	options := getQuizOptions(optionValues)
	quiz := models.Quiz{
		Id:              id,
		TotalQuestions:  1,
		Question: r.questioner.GetQuestion(),
		Players: make(map[*websocket.Conn]*models.Player),
		Options: options,
	}

	quizzes[id] = &quiz
	log.Info(`Created quiz`, id)
}

func (r *quizMaster) Broadcast(id string, message interface{}) {
	for key, _ := range quizzes[id].Players {
		_ = key.WriteJSON(message)
	}
}

func (r *quizMaster) CheckAnswer(id string, connection *websocket.Conn, message *models.Message) {
	currentQuiz := quizzes[id]
	log.Printf("|%s|%s|", currentQuiz.Question.CorrectAnswer, strings.Trim(message.Answer, ` `))
	if currentQuiz.Question.CorrectAnswer == strings.Trim(message.Answer, ` `) {
		currentQuiz.Players[connection].Score++
		currentQuiz.Question = r.questioner.GetQuestion()

		state := getQuizState(currentQuiz)

		rs := models.NewMessage()
		rs.Type = "correct-answer"
		_ = connection.WriteJSON(rs)

		r.Broadcast(id, state)
	} else {
		log.Println("WRONG ANSWER RS");

		rs := models.NewMessage()
		rs.Type = "wrong-answer";

		// send wrong response to
		err := connection.WriteJSON(rs)
		if err != nil {
			log.Println(err, "writing error")
		}
	}
}

func (r *quizMaster) AddPlayer(id string, connection *websocket.Conn, player *models.Player) {
	quizzes[id].Players[connection] = player
	r.sendQuizState(quizzes[id])
}

func (r *quizMaster) DeletePlayer(id string, connection *websocket.Conn) {
	delete(quizzes[id].Players, connection)
	_ = connection.Close()
}

func (r *quizMaster) QuizExists(id string) bool {
	return quizzes[id] != nil
}

func getQuizState(quiz *models.Quiz) *models.QuizState {
	quizState := models.NewQuizState()
	quizState.CurrentQuestion = quiz.Question
	quizState.TotalQuestions = quiz.TotalQuestions

	for _, val := range quiz.Players {
		quizState.Players = append(quizState.Players, val)
	}

	return quizState
}

func (r *quizMaster) sendQuizState(quiz *models.Quiz) {
	quizState := getQuizState(quiz)
	r.Broadcast(quiz.Id, quizState)
}

func getQuizOptions(paramString url.Values) map[string]bool {
	ret := make(map[string]bool)

	optionValues := paramString["options"]
	if (len(optionValues) == 0) {
		return ret;
	}

	err := json.Unmarshal([]byte(optionValues[0]), &ret)
	if err != nil {
		log.Error("Quiz options could not be parsed")
	}

	return ret
}
