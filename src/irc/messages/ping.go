package irc

type m_ping struct {
	Message string
}

const ping_tmpl = "PONG %s"

func NewPingMessage(message string) m_ping {
	return m_ping{message}
}

func (m m_ping) Tmpl() string {
	return ping_tmpl
}

func (m m_ping) Data() []interface{} {
	return []interface{}{m.Message}
}
