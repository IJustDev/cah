package src

type PlayerReceivedAnswersEventPayload struct {
	Player  *Player
	Answers []Answer
}

type playerReceivedAnswersEvent struct {
	handlers []func(payload PlayerReceivedAnswersEventPayload)
}

var PlayerReceivedAnswersEvent playerReceivedAnswersEvent

func (e *playerReceivedAnswersEvent) Register(handler func(payload PlayerReceivedAnswersEventPayload)) {
	e.handlers = append(e.handlers, handler)
}

func (e *playerReceivedAnswersEvent) Trigger(payload PlayerReceivedAnswersEventPayload) {
	for _, handler := range e.handlers {
		go handler(payload)
	}
}
