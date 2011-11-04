package irc

type m_private struct {
	target string
	message string
}

func NewPrivateMessage(target, message string) m_private {
	return m_private{target, message}
}

func (m m_private) String() string {
	return ":source PRIVMSG " + m.target + " :" + m.message + "\n"
}
