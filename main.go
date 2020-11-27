package main

import (
	"arabic-server/models"
	"arabic-server/services"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
};

var quizzes = map[string]*models.Quiz{}
var clients = map[*websocket.Conn]string{} // quiz ID is the value
var questioner services.Questioner

func init() {
	questioner = services.NewQuestioner()
}

func home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := uuid.NewV4().String()
	players := make(map[*websocket.Conn]*models.Player);
	quiz := models.Quiz{
		Id:              id,
		TotalQuestions:  1,
		Question: questioner.GetQuestion(),
		Players:      players,
	}

	quizzes[id] = &quiz

	rs := fmt.Sprintf(`{"status": "OK", "quiz-id":"%s" }`, id)
	w.Write([]byte(rs))
}

func createQuizState(quiz *models.Quiz) *models.QuizState {
	quizState := models.NewQuizState()
	quizState.CurrentQuestion = quiz.Question
	quizState.TotalQuestions = quiz.TotalQuestions

	for _, val := range quiz.Players {
		quizState.Players = append(quizState.Players, val)
	}

	return quizState;
}

func broadcast(players map[*websocket.Conn]*models.Player, data interface{}) {
	for key, _ := range players {
		_ = key.WriteJSON(data)
	}
}

func Quiz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println(params.ByName("quiz"));

	quizId := params.ByName("quiz")
	name := params.ByName("name")

	c, err := upgrader.Upgrade(w, r, nil)
	if (err != nil) {
		log.Println("upgrade: ", err)
		return
	}

 	defer func(quiz *models.Quiz) {
		delete(quizzes[quizId].Players, c)
		c.Close()
	}(quizzes[quizId])


	// Add new player to the collection
	player := models.Player {
		Name:name,
		Score: 0,
	}
	quizzes[quizId].Players[c] = &player

	// Send current state to new player
	currentQuiz := quizzes[quizId]
	quizState := createQuizState(currentQuiz)
	for key, _ := range currentQuiz.Players {
		err = key.WriteJSON(quizState)
		if err != nil {
			log.Println("FAILED TO SEND QUIZ STATE TO ", err, name, quizId)
		}
	}

	for {
		message := models.NewMessage()
		err = c.ReadJSON(message)
		if err != nil {
			log.Println("read: ", err)
			break
		}


		log.Printf("recv:", message)


		if message.Type == "answer" {
			log.Printf("|%s|%s|", currentQuiz.Question.CorrectAnswer, strings.Trim(message.Answer, ` `))
			if currentQuiz.Question.CorrectAnswer == strings.Trim(message.Answer, ` `) {
				currentQuiz.Players[c].Score++
				currentQuiz.Question = questioner.GetQuestion()

				state := createQuizState(currentQuiz)

				rs := models.NewMessage()
				rs.Type = "correct-answer"
				_ = c.WriteJSON(rs)

				broadcast(currentQuiz.Players, state)
			} else {
				log.Println("WRONG ANSWER RS");

				rs := models.NewMessage()
				rs.Type = "wrong-answer";

				// send wrong response to
				err = c.WriteJSON(rs)
				if err != nil {
					log.Println(err, "writing error")
				}
			}
		}
	}
}



func main() {
	router := httprouter.New()
	router.POST("/", home)
	router.GET("/quizzes/:quiz/:name", Quiz)

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
}
