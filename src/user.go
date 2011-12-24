package main

type User struct {
	ircConns []Irc.Connection
	webConns []Web.Connection
	username string
}
