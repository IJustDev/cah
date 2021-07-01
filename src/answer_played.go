package src

var AnswerPlayedEvent answerPlayedEvent

type answerPlayedEvent struct {
	handlers []func(playerAnswer PlayerAnswer, g Game)
}

func (e *answerPlayedEvent) Register(handler func(playerAnswer PlayerAnswer, g Game)) {
	e.handlers = append(e.handlers, handler)
}

func (e *answerPlayedEvent) Trigger(playerAnswer PlayerAnswer, g Game) {
	for _, handler := range e.handlers {
		go handler(playerAnswer, g)
	}
}
