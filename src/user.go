package main

// This is the basic layout. it won't compile currently.
type User struct {
	connections []Irc.Connection
	comm        Web.Connection
	username    string
}
