package irc

const tmpl_part = "PART %s :%s\n"

type m_part struct {
	channels string
	reason string
}

func NewPartMessage(rooms []string, reason string) m_part {
	chans := ""
	for i, val := range rooms {
		if i > 0 {
			chans += ","
		}
		chans += val
	}
	if reason == "" {
		reason = "Leaving."
	}
	return m_part{chans, reason}
}

func (m m_part) Tmpl() string {
	return tmpl_part
}

func (m m_part) Data() []interface{} {
	return []interface{}{m.channels, m.reason}
}
