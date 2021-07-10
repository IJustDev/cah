package src

type gameCreatedEvent struct {
	handlers []func(g Game)
}

var GameCreatedEvent gameCreatedEvent

func (e *gameCreatedEvent) Register(handler func(g Game)) {
	e.handlers = append(e.handlers, handler)
}

func (e *gameCreatedEvent) Trigger(g Game) {
	for _, handler := range e.handlers {
		go handler(g)
	}
}
