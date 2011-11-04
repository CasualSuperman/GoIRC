package irc

type m_quit struct {
	reason string
}

func NewQuitMessage(reason string) m_quit {
	return m_quit{reason}
}

func (m m_quit) String() string {
	if m.reason == "" {
		return ":source QUIT :Leaving.\n"
	}
	return ":source QUIT :" + m.reason + "\n"
}
