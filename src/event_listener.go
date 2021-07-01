package src

type EventListener struct {
}

func NewEventListener() *EventListener {
	return &EventListener{}
}

func (e *EventListener) Subscribe() {
	GameStartedEvent.Register(gameStarted)
	RoundStartedEvent.Register(roundStarted)
	PlayerJoinedEvent.Register(playerJoined)
	AnswerPlayedEvent.Register(answerPlayed)
	AnswerPickedEvent.Register(answerPicked)
	PlayerReceivedAnswersEvent.Register(playerReceivedAnswers)
}

func gameStarted(g Game) {
}

func roundStarted(payload RoundStartedEventPayload) {

}

func playerJoined(payload PlayerJoinedEventPayload) {

}

func answerPlayed(playerAnswer PlayerAnswer, g Game) {

}

func answerPicked(playerAnswer PlayerAnswer, g Game) {

}

func answersRevealed(playerAnswer []PlayerAnswer, g Game) {

}

func playerReceivedAnswers(payload PlayerReceivedAnswersEventPayload) {

}
