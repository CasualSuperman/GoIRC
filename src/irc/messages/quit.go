package irc

type m_quit struct {
	reason string
}

func NewQuitMessage(reason string) {
	return m_quit{reason}
}

func (m m_quit) String() string {
	if m.reason == "" {
		return ":source QUIT :Leaving."
	}
	return ":source QUIT :" + m.reason
}
