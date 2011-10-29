package main

import strings

type m_part struct {
	channels []string
	reason string
}

func NewPartMessage(rooms []string, reason string) {
	return m_part{append(nil, rooms), reason}
}

func (m m_part) String() string {
	s := ":source PART "
	for i, val := range m.channels {
		if i > 0 {
			s += ","
		}
		if strings.TrimL(val, "#") == val {
			s += "#" + val
		} else {
			s += val
		}
	}
	s += " :"
	s += m.reason
	return s
}
