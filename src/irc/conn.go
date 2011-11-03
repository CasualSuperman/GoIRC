package irc

import (
	"bufio"
	"net"
	"os"
)

type Conn struct {
	serverConn net.Conn
	Recv <-chan message
	Send chan<- message
}

func Connect(server string) (c Conn, err os.Error) {
	stream, err := net.Dial("tcp", server)
	if err != nil {
		return c, err
	}
	recv := make(chan message, 5)
	send := make(chan message, 5)

	c = Conn{stream,
			 recv,
			 send}
	go handle(stream, recv, send)
	return c, nil
}

func handle(conn net.Conn, recv, send chan message) {
	io := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	data := make(chan message)
	go process(io, data)
	for {
		select {
			case i := <-send:
				io.WriteString(i.String() + "\n")
				io.Flush()
			case recv <- <- data:
		}
	}
}
