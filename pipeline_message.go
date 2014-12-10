package main

import "strings"

type PipelineMessage struct {
	Model  string
	Mction string
}

func NewPipelineMessage(msg string) *PipelineMessage {
	s := strings.Split(msg, ":")
	action, model := s[0], s[1]
	return &PipelineMessage{model, action}
}
