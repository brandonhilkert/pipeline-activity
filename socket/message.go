package socket

type Message struct {
	Model  string `json:"model"`
	Action string `json:"action"`
}

func (m *Message) String() string {
	return m.Model + " was " + m.Action + "d"
}
