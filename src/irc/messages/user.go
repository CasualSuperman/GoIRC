package irc

type m_user struct {
	username string
	hidden bool
	realName string
}

func NewUserMessage(username string, hidden bool, realName string) m_user {
	m := m_user{username, hidden, realName}
	return m
}

func (m m_user) String() string {
	s := "USER"
	s += " "
	s += m.username
	s += " "
	if m.hidden {
		s += "8"
	} else {
		s += "0"
	}
	s += " "
	s += "*"
	s += " "
	s += ":"
	s += m.realName
	return s
}
