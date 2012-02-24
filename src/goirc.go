package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	encoder := json.NewEncoder(ws)
	i := 0
	for _, line := range lines {
		i++
		if (i > 47) {
			time.Sleep(time.Duration((rand.Intn(2000) + 700) * int(time.Millisecond)))
		}
		data := struct{Server, Line string}{"irc.foonetic.net",line}
		encoder.Encode(data)
	}
	fmt.Println("Conn closed.")
	ws.Close()
}
