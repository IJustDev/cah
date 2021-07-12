package src

type Context struct {
	Broadcast chan []byte
}

func NewContext(broadcast chan []byte) *Context {
	return &Context{
		Broadcast: broadcast,
	}
}
