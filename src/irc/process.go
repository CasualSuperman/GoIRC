package irc

import (
	"bufio"
	"net"
)

func process(read *bufio.ReadWriter, to chan message) {
		for {
			// Try to read from the network
			line, isPrefix, err := read.ReadLine()
			if err != nil {
				nerr, ok := err.(net.Error)
				if ok && !nerr.Timeout() {
					panic("ERROR: " + nerr.String())
				} else if !ok {
					panic("ERROR: " + err.String())
				}
			}
			buff := ""
			for isPrefix {
				buff += string(line)
				line, isPrefix, err = read.ReadLine()
			}
			buff += string(line)
			to <- NewMessage(buff)
		}
}
