package src

var AnswerPickedEvent answerPickedEvent

type answerPickedEvent struct {
	handlers []func(playerAnswer PlayerAnswer, g Game)
}

func (e *answerPickedEvent) Register(handler func(playerAnswer PlayerAnswer, g Game)) {
	e.handlers = append(e.handlers, handler)
}

func (e *answerPickedEvent) Trigger(playerAnswer PlayerAnswer, g Game) {
	for _, handler := range e.handlers {
		go handler(playerAnswer, g)
	}
}
