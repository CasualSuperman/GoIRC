package irc

const tmpl_nick = "NICK %s\n"

type m_nick struct {
	nick string
}

func NewNickMessage(nick string) m_nick {
	m := m_nick{nick}
	return m
}

func (m m_nick) Tmpl() string {
	return tmpl_nick
}

func (m m_nick) Data() []interface{} {
	return []interface{}{m.nick}
}
