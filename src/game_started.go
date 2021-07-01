package src

type gameStartedEvent struct {
	handlers []func(g Game)
}

var GameStartedEvent gameStartedEvent

func (e *gameStartedEvent) Register(handler func(g Game)) {
	e.handlers = append(e.handlers, handler)
}

func (e *gameStartedEvent) Trigger(g Game) {
	for _, handler := range e.handlers {
		go handler(g)
	}
}
