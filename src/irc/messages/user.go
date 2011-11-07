package irc

const tmpl_user = "USER %s %d * :%s\n"

type m_user struct {
	username string
	hidden bool
	realName string
}

func NewUserMessage(username string, hidden bool, realName string) m_user {
	m := m_user{username, hidden, realName}
	return m
}

func (m m_user) Tmpl() string {
	return tmpl_user
}

func (m m_user) Data() []interface{} {
	i := 0
	if m.hidden {
		i = 8
	}
	return []interface{}{m.username, i, m.realName}
}
