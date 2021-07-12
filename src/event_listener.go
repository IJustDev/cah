package src

import (
	"encoding/json"
)

type EventListener struct {
}

func NewEventListener() *EventListener {
	return &EventListener{}
}

func (e *EventListener) Subscribe(broadcast chan []byte) {
	context := NewContext(broadcast)
	GameCreatedEvent.Register(gameCreated(context))
	RoundStartedEvent.Register(roundStarted(context))
	PlayerJoinedEvent.Register(playerJoined(context))
	AnswerPlayedEvent.Register(answerPlayed(context))
	AnswerPickedEvent.Register(answerPicked(context))
	PlayerReceivedAnswersEvent.Register(playerReceivedAnswers(context))
}

func gameCreated(ctx *Context) func(g Game) {
	type broadcastResponse struct {
		Type string
		Id   string
	}
	return func(g Game) {
		response := &broadcastResponse{
			Type: "game_created_event",
			Id:   g.Id,
		}

		json, _ := json.Marshal(response)

		ctx.Broadcast <- []byte(json)
	}
}

func roundStarted(ctx *Context) func(payload RoundStartedEventPayload) {
	type broadcastResponse struct {
		Type       string
		GameId     string
		Round      Round
		FirstRound bool
	}
	return func(payload RoundStartedEventPayload) {
		response := broadcastResponse{
			Type:       "round_started_event",
			GameId:     payload.Game.Id,
			Round:      payload.Round,
			FirstRound: true,
		}

		json, _ := json.Marshal(response)
		ctx.Broadcast <- []byte(json)
	}

}

func playerJoined(ctx *Context) func(payload PlayerJoinedEventPayload) {
	return func(payload PlayerJoinedEventPayload) {
	}

}

func answerPlayed(ctx *Context) func(playerAnswer PlayerAnswer, g Game) {
	return func(playerAnswer PlayerAnswer, g Game) {
	}

}

func answerPicked(ctx *Context) func(playerAnswer PlayerAnswer, g Game) {
	return func(playerAnswer PlayerAnswer, g Game) {
	}

}

func answersRevealed(ctx *Context) func(playerAnswer []PlayerAnswer, g Game) {
	return func(playerAnswer []PlayerAnswer, g Game) {
	}

}

func playerReceivedAnswers(ctx *Context) func(payload PlayerReceivedAnswersEventPayload) {
	return func(payload PlayerReceivedAnswersEventPayload) {
	}

}
