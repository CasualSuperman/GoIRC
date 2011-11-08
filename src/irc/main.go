package main

import (
	irc "./_obj/irc"
	"fmt"
)

func main() {
	user := irc.NewUser("goirc_tester", "rwertman", false, "Robert Wertman")
	conn, err := irc.Connect("irc.foonetic.net:6667", user)
	if err != nil {
		panic(err)
	}
	i := 0
	ok := true
	for ok {
		i++
		var data irc.Message
		data, ok = <-conn.Recv()
		fmt.Printf(data.Tmpl(), data.Data()...)
		if i == 20 {
			conn.Send(irc.NewJoinMessage("#ufeff"))
		}
		if i > 200 {
			ok = false
		}
	}
}
