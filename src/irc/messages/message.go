package irc

import "os"

type Message interface {
	String() string
}

func NewMessage(in string) Message {
	return os.NewError(in)
}
