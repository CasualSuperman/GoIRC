package irc

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func process(read *bufio.ReadWriter, to chan Message) {
		for {
			// Try to read from the network
			line, isPrefix, err := read.ReadLine()
			if err != nil {
				nerr, ok := err.(net.Error)
				if ok && !nerr.Timeout() {
					to <- NewMessage("ERROR: " + nerr.String())
				} else if !ok {
					to <- NewMessage("ERROR: " + err.String())
				}
			}
			buff := ""
			for isPrefix {
				buff += string(line)
				line, isPrefix, err = read.ReadLine()
			}
			buff += string(line)
			if strings.Index(buff, "PING") == 0 {
				resp := strings.Replace(buff, "PING", "PONG", 1)
				read.WriteString(resp + "\n")
				read.Flush()
				fmt.Println(buff + " (" + resp + ")")
			} else {
				to <- NewMessage(buff)
			}
		}
}
