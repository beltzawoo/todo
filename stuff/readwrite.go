package stuff

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
)

type todo []string

func getPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.todo.json"
}

func GetTodo() todo {
	data, err := ioutil.ReadFile(getPath())
	if err != nil {
		log.Fatal(err)
	}

	var todoSlice todo

	err = json.Unmarshal(data, &todoSlice)
	if err != nil {
		log.Fatal(err)
	}
	return todoSlice
}

func writeTodo(s []string) {
	b, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(getPath(), b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
