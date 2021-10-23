package controllers

import (
	"arabic-server/services"
	"fmt"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type HomeController interface {
	Index(w http.ResponseWriter, req *http.Request, params httprouter.Params)
}

func NewHomeController(quizMaster services.QuizMaster) HomeController {
	return &homeController{
		quizMaster: quizMaster,
	}
}

type homeController struct {
	quizMaster services.QuizMaster
}

func (r *homeController) Index(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := uuid.NewV4().String()
	r.quizMaster.CreateQuiz(id, req.URL.Query())
	rs := fmt.Sprintf(`{"status": 200, "quiz-id":"%s" }`, id)
	_, _ = w.Write([]byte(rs))
}
