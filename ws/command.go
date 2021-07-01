package ws

type Command interface {
	Name() string
	ParameterInt() int
	Parameters()
	CommandResult()
}

type CommandResult interface {
}
