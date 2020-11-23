package main

import (
	"os"
	"strconv"
	"todo/stuff"
)

func main() {
	if len(os.Args) < 2 {
		stuff.PrettyPrint(stuff.GetTodo())
	} else {
		switch os.Args[1] {
		case "a":
			stuff.AppendTodo(os.Args[2])
		case "r":
			stuff.RemoveTodo(strconv.Atoi(os.Args[2]))
		}
	}

}
