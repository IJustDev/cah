package src

type RoundStartedEventPayload struct {
	Round      Round
	Game       Game
	RoundCount int
}

type roundStartedEvent struct {
	handlers []func(RoundStartedEventPayload)
}

var RoundStartedEvent roundStartedEvent

func (e *roundStartedEvent) Register(handler func(RoundStartedEventPayload)) {
	e.handlers = append(e.handlers, handler)
}

func (e *roundStartedEvent) Trigger(payload RoundStartedEventPayload) {
	for _, handler := range e.handlers {
		go handler(payload)
	}
}
