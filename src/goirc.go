package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"websocket"
)

func main() {
	http.Handle("/", websocket.Handler(Handle))
	fmt.Println("Handling.")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed.")
		fmt.Println(err)
	}
}

func Handle(ws *websocket.Conn) {
	data := make(map[string]string)
	done := false

	data["msg"] = "Hello World!"

	fmt.Printf("New Connection from %s.\n", ws.RemoteAddr())

	reader := json.NewDecoder(ws)
	writer := json.NewEncoder(ws)

	writer.Encode(data)

	for !done {
		reader.Decode(&data)

		if data["action"] == "disconnect" {
			done = true
		} else {
			fmt.Println(data)
		}

		data := make(map[string]string)
	}
	fmt.Println("Conn closed.")
	ws.Close()
}
