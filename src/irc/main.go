package main

import (
	irc "./_obj/irc"
	"fmt"
	"os/signal"
)

func main() {
	user := irc.NewUser("goirc_tester", "rwertman", false, "Robert Wertman")
	conn, err := irc.Connect("irc.foonetic.net:6667", user)
	if err != nil {
		panic(err)
	}
	i := 0
	done := false
	quit := make(chan bool)
	go func(quit chan bool) {
		for !done {
			data := <-signal.Incoming
			if data.String() == "SIGINT: interrupt" {
				quit <- true
			} else {
				fmt.Println(data.String())
			}
		}
	}(quit)
	for !done {
		i++
		select {
			case data, ok := <-conn.Recv():
				fmt.Printf(data.Tmpl(), data.Data()...)
				if !ok {
					done = true
				}
			case done = <-quit:
		}
		if i == 20 {
			conn.Send(irc.NewJoinMessage("#ufeff"))
		}
	}
	conn.Send(irc.NewQuitMessage("Closed"))
}
