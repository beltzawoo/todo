package stuff

func AppendTodo(n string) {
	s := append(GetTodo(), n)
	writeTodo(s)
}

func RemoveTodo(n int, err error) {
	t := GetTodo()
	s := append(t[:n-1], t[n:]...)
	writeTodo(s)
}
