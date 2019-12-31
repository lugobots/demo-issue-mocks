package bar

type DoShomething interface {
	DoIt() bool
	DoThat() string
}

type MoreStuff interface {
	TryToDo(DoShomething) bool
	AskTodo() DoShomething
}
