package main

import (
	irc "./conn"
	"fmt"
)

func main() {
	conn, err := irc.Connect("irc.foonetic.net:6667")
	if err != nil {
		panic(err)
	}
	i := 0
	ok := true
	for ok {
		i++
		var data irc.Message
		data, ok = <-conn.Recv
		fmt.Println(data.String())
		if i == 2 {
			fmt.Println("NICK CasualSuperman")
			conn.ServerConn.Write([]byte("NICK CasualSuperman\n"))
			fmt.Println("USER rwertman 8 * :Robert Wertman")
			conn.ServerConn.Write([]byte("USER rwertman 8 * :Robert Wertman\n"))
		}
		if i > 100 {
			break
		}
	}
}
