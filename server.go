package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Our database
var todoList []string

//Helper functions to parse the input between a message and the actual task
func getCmd(input string) string {
	inputArr := strings.Split(input, " ")
	return inputArr[0]
}

func getMessage(input string) string {
	inputArr := strings.Split(input, " ")
	var result string

	for i := 1; i < len(inputArr); i++ {
		result += inputArr[i]
	}

	return result
}

//end of helpers

//Adding a task to the todolist array.
func updateTodoList(input string) {
	tmpList := todoList
	todoList = []string{}

	for _, val := range tmpList {

		if val == input {
			continue
		}

		todoList = append(todoList, val)
	}
}

func main() {

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}

		defer conn.Close()

		for {

			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read failed", err)
				break
			}

			//input
			input := string(message)

			cmd := getCmd(input)
			msg := getMessage(input)

			if cmd == "add" {
				todoList = append(todoList, msg)
			} else if cmd == "done" {
				updateTodoList(msg)
			}

			output := "Current Todos: \n"
			for _, todo := range todoList {
				output += "\n - " + todo + "\n"
			}

			output += "\n----------------------------------------"

			message = []byte(output)
			err = conn.WriteMessage(mt, message)

			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	})

	//pointing to our html file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
