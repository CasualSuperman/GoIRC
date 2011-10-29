package main

import "strings"

type m_join struct {
	channel string
}

func (m m_join) String() string {
	s := ":source JOIN :"
	if strings.TrimLeft(m.channel, "#") == m.channel {
		s += "#" + m.channel
	} else {
		s += m.channel
	}
	return s
}
