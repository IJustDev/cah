package src

type PlayerJoinedEventPayload struct {
	Player Player
	Game   Game
}

type playerJoinedEvent struct {
	handlers []func(PlayerJoinedEventPayload)
}

var PlayerJoinedEvent playerJoinedEvent

func (e *playerJoinedEvent) Register(handler func(PlayerJoinedEventPayload)) {
	e.handlers = append(e.handlers, handler)
}

func (e *playerJoinedEvent) Trigger(payload PlayerJoinedEventPayload) {
	for _, handler := range e.handlers {
		go handler(payload)
	}
}
