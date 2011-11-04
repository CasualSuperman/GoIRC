package irc

type user struct {
	Nick, Username, Name string
	Hidden bool
}

func NewUser(nick, username string, hidden bool, name string) user {
	return user{nick, username, name, hidden}
}

func (u user) UserMessage() Message {
	return NewUserMessage(u.Username, u.Hidden, u.Name)
}

func (u user) NickMessage() Message {
	return NewNickMessage(u.Nick)
}
