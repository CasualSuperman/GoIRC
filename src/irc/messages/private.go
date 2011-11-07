package irc

const tmpl_private = "PRIVMSG %s :%s\n"

type m_private struct {
	target string
	message string
}

func NewPrivateMessage(target, message string) m_private {
	return m_private{target, message}
}

func (m m_private) Tmpl() string {
	return tmpl_private
}

func (m m_private) Data() []interface{} {
	return []interface{}{m.target, m.message}
}
