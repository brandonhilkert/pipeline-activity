package pubsub

import "strings"

type Message struct {
	Model  string
	Action string
}

func (m *Message) String() string {
	return "Model: " + m.Model + " Action: " + m.Action
}

func NewMessage(msg string) *Message {
	s := strings.Split(msg, ":")
	action, model := s[0], s[1]
	return &Message{model, action}
}
