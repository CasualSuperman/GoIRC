package irc

type m_nick struct {
	nick string
}

func NewNickMessage(nick string) m_nick {
	m := m_nick{nick}
	return m
}

func (m m_nick) String() string {
	s := "NICK"
	s += " "
	s += m.nick
	s += "\n"
	return s
}
