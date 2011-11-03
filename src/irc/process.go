package irc

import (
	"bufio"
	"net"
)

func process(read *bufio.ReadWriter, to chan message) {
		defer func() {
			recover()
			process(read, to)
		}()
		for {
			// Try to read from the network
			line, isPrefix, err := read.ReadLine()
			if err != nil {
				nerr := err.(net.Error)
				if !nerr.Timeout() {
					panic("ERROR:" + nerr.String())
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
