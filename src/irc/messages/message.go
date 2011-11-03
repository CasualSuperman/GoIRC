package irc

type message interface {
	String() string
}

func NewMessage(in string) message {
	return nil
}
