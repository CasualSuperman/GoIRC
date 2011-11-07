package main

import (
	irc "./_obj/irc"
	"fmt"
)

func main() {
	user := irc.NewUser("CasualSuperman", "rwertman", false, "Robert Wertman")
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
		if i > 100 {
			ok = false
		}
	}
}
