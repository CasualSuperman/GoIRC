package irc

type err struct {
	err string
}

func (e err) String() string {
	return e.err
}

func error(s string) err {
	return err{s}
}
