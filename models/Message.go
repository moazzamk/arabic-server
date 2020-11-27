package models

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Answer string `json:"answer"`
}

func NewMessage() *Message {
	return &Message{}
}