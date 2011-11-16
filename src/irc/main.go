package main

import (
	irc "./_obj/irc"
	"bufio"
	"fmt"
	"os"
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
	input := make(chan string)
	for !done {
		i++
		select {
			case data, ok := <-conn.Recv():
				fmt.Printf(data.Tmpl(), data.Data()...)
				if !ok {
					done = true
				}
			case send := <-input:
				conn.Send(irc.NewPrivateMessage("#ufeff", send))
			case done = <-quit:
		}
		if i == 20 {
			conn.Send(irc.NewJoinMessage("#ufeff"))
			go func() {
				line := ""
				reader := bufio.NewReader(os.Stdin)
				for {
					var buf []byte
					isPrefix := true
					for isPrefix {
						buf, isPrefix, _ = reader.ReadLine()
						line += string(buf)
					}
					input <- line
					line = ""
				}
			}()
		}
	}
	conn.Send(irc.NewQuitMessage("Closed"))
}
