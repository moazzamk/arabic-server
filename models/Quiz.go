package models

import "github.com/gorilla/websocket"

type Quiz struct {
	Id string `json:"quiz-id"`
	TotalQuestions int `json:"total-questions"`
	Question *Question `json:"question"`
	Players map[*websocket.Conn]*Player
}
