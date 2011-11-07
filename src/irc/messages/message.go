package irc

type Message interface {
	Tmpl() string
	Data() []interface{}
}
