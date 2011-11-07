package irc

type m_quit struct {
	reason string
}

const tmpl_quit = "QUIT :%s\n"

func NewQuitMessage(reason string) m_quit {
    if reason == "" {
		reason = "Leaving."
	}
	return m_quit{reason}
}

func (m m_quit) Tmpl() string {
	return tmpl_quit
}

func (m m_quit) Data() []interface{} {
	return []interface{}{m.reason}
}
