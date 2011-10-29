package main

import strings

type m_join struct {
	channel string
}

func (m m_join) String() string {
	if strings.TrimL(m.channel, "#") == m.channel {
		return ":source JOIN :#" + m.channel
	} else {
		return ":source JOIN :" + m/channel
	}
}
