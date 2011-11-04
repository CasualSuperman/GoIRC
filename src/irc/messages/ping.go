package irc

type m_ping struct {
	message string
}

func NewPingMessage(message string) m_ping {
	return m_ping{message}
}

func (m m_ping) String() string {
	return "PING :" + m.message + "\n"
}
