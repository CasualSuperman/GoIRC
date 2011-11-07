package irc

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Conn struct {
	ServerConn net.Conn
	info user
	recv *chan Message
	send *chan Message
}

func Connect(server string, info user) (c Conn, err os.Error) {
	stream, err := net.Dial("tcp", server)
	if err != nil {
		return c, err
	}
	recv := make(chan Message, 5)
	send := make(chan Message, 5)
	go func() {
		send <- info.NickMessage()
		send <- info.UserMessage()
	}()
	c = Conn{stream,
			 info,
			 &recv,
			 &send}
	go handle(stream, &recv, &send)
	return c, nil
}

func (c Conn) Recv() <-chan Message {
	return *c.recv
}

func (c Conn) Send() <-chan Message {
	return *c.send
}

func handle(conn net.Conn, recv, send *chan Message) {
	io := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	data := make(chan Message)
	go process(io, data)
	for {
		select {
			case i := <-*send:
				fmt.Println("SENDING: " + i.String())
				//io.WriteString(i.String() + "\n")
				conn.Write([]byte(i.String()))
				//io.Flush()
			case *recv <- <- data:
		}
	}
}
