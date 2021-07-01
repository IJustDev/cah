package events

/**
type gameStartedEvent struct {
	handlers []interface{ Handle(GameStartedEventPayload) }
}

var GameStartedEvent gameStartedEvent

type GameStartedEventPayload struct {
	Id      int
	Players []Player
}

func (e *gameStartedEvent) Register(handler interface{ Handle(GameStartedEventPayload) }) {
	e.handlers = append(e.Handlers, handler)
}

func (e *gameStartedEvent) Trigger(payload gameStartedEventPayload) {
	for _, handler := range e.handlers {
		go handler.Handle(payload)
	}
}
*/
