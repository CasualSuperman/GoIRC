package main

import (
	"bufio"
	"fmt"
	"http"
	"websocket"
)

func main() {
	http.Handle("/irc", websocket.Handler(Handle))
	fmt.Println("Handling.")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed.")
	}
}

func Handle(ws *websocket.Conn) {
	fmt.Println("New Connection.")
	fmt.Println(ws.Protocol)
	fmt.Println(ws.Request)
	reader := bufio.NewReader(ws)
	ws.Write([]uint8("Hello, world.\n"))
	for {
		fmt.Println(reader.ReadLine())
	}
	fmt.Println("Conn closed.")
}
