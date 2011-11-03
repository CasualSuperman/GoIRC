package irc

import (
	"net"
	"os"
)

type Conn struct {
	serverConn net.Conn
}

func Connect(server string) Conn, os.Error {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		return nil, err
	}
	return Conn{conn}
}
