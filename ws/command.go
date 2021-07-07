package ws

type Command interface {
	Command() string
	Parameters() []string
	Execute(map[string]string) string
}
