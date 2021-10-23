package controllers

import (
	"arabic-server/models"
	"arabic-server/services"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)


type QuizController interface {
	Quiz(w http.ResponseWriter, req *http.Request, params httprouter.Params)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// NewQuizController returns a new instance of the quiz controller
func NewQuizController(quizMaster services.QuizMaster) QuizController {
	return &quizController{
		quizMaster: quizMaster,
	}
}

type quizController struct {
	quizMaster services.QuizMaster
}

func (r *quizController) Quiz(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	quizId := strings.Trim(params.ByName("quiz"), ` `)
	name := strings.Trim(params.ByName("name"), ` `)

	if quizId == `` || name == `` {
		_, _ = w.Write([]byte(`{"status": 400, "message":"quiz and name are required"}`))
		return
	}

	if !r.quizMaster.QuizExists(quizId) {
		r.quizMaster.CreateQuiz(quizId, req.URL.Query())
	}
	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}

	defer func(qm services.QuizMaster, quizId string) {
		qm.DeletePlayer(quizId, c)
		_ = c.Close()
	}(r.quizMaster, quizId)

	// Add new player to the collection
	player := models.Player {
		Name:name,
		Score: 0,
	}
	r.quizMaster.AddPlayer(quizId, c, &player)

	// Listen to messages on the socket
	for {
		message := models.NewMessage()
		err = c.ReadJSON(message)
		if err != nil {
			log.Println("read: ", err)
			break
		}

		log.Printf("recv: %v", message)
		if message.Type == "answer" {
			r.quizMaster.CheckAnswer(quizId, c, message)
		}
	}
}

