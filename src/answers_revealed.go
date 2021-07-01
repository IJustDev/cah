package src

var AnswersRevealedEvent answersRevealedEvent

type answersRevealedEvent struct {
	handlers []func(playerAnswer []PlayerAnswer, g Game)
}

func (e *answersRevealedEvent) Register(handler func(playerAnswer []PlayerAnswer, g Game)) {
	e.handlers = append(e.handlers, handler)
}

func (e *answersRevealedEvent) Trigger(playerAnswer []PlayerAnswer, g Game) {
	for _, handler := range e.handlers {
		go handler(playerAnswer, g)
	}
}
