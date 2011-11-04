package irc

import "strings"

type m_join struct {
	channel string
}

func NewJoinMessage(channel string) m_join{
	return m_join{channel}
}

func (m m_join) String() string {
	if m.channel == "" {
		return ":source JOIN 0\n"
	}
	s := ":source JOIN :"
	if strings.TrimLeft(m.channel, "#") == m.channel {
		s += "#" + m.channel
	} else {
		s += m.channel
	}
	s += "\n"
	return s
}
