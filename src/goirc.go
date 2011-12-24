package main

import (
	//"bufio"
	"fmt"
	"http"
	"io/ioutil"
	"json"
	"strings"
	"websocket"
)

var lines []string

func main() {
	data, err := ioutil.ReadFile("log")
	str := string(data)
	lines = strings.Split(str, "\n")

	http.Handle("/", websocket.Handler(Handle))
	fmt.Println("Handling.")
	err = http.ListenAndServe(":3654", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed.")
		fmt.Println(err)
	}
}

func Handle(ws *websocket.Conn) {
	fmt.Println("New Connection.")
	fmt.Println(ws.Protocol)
	fmt.Println(ws.Request)
	encoder := json.NewEncoder(ws)
	for _, line := range lines {
		encoder.Encode(line)
	}
	fmt.Println("Conn closed.")
	ws.Close()
}
