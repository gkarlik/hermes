package hermes

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	REGUSER = iota
	UNREGUSER
	MESSAGE
	BROADCAST
)

type Message struct {
	Type      int
	Recipient string
	Sender    string
	Body      string
}

func NewMessage(messageType int, sender, recipient, body string) *Message {
	return &Message{
		Type:      messageType,
		Recipient: recipient,
		Sender:    sender,
		Body:      body,
	}
}

func ParseMessage(msg []string) (*Message, error) {
	if len(msg) != 4 {
		return nil, errors.New("Invalid message")
	}

	t, _ := strconv.Atoi(msg[0])
	message := &Message{
		Type:      t,
		Sender:    msg[1],
		Recipient: msg[2],
		Body:      msg[3],
	}

	return message, nil
}

func (m *Message) ToStringArray() []string {
	return []string{
		fmt.Sprintf("%d", m.Type),
		m.Sender,
		m.Recipient,
		m.Body,
	}
}
