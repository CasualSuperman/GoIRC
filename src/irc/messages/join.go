package irc

const tmpl_join = "JOIN %s\n"

type m_join struct {
	channel string
}

func NewJoinMessage(channel string) m_join {
	if channel == "" {
		channel = "0"
	} else {
		channel = ":" + channel
	}
	return m_join{channel}
}

func (m m_join) Data() []interface{} {
	return []interface{}{m.channel}
}

func (m m_join) Tmpl() string {
	return tmpl_join
}
