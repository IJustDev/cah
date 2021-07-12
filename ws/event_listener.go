package ws

import (
	"encoding/json"

	cah "github.com/royalzsoftware/cah/src"
)

type EventListener struct {
}

const (
	ALL = iota
	ONLY_GAME_MEMBERS
)

type Notification struct {
	Type    string
	Scope   int
	GameId  string
	Payload interface{}
}

func NewEventListener() *EventListener {
	return &EventListener{}
}

func (e *EventListener) Subscribe(broadcast chan []byte) {
	context := NewContext(broadcast)
	cah.GameCreatedEvent.Register(gameCreated(context))
	cah.RoundStartedEvent.Register(roundStarted(context))
	cah.PlayerJoinedEvent.Register(playerJoined(context))
	cah.AnswerPlayedEvent.Register(answerPlayed(context))
	cah.AnswerPickedEvent.Register(answerPicked(context))
	cah.PlayerReceivedAnswersEvent.Register(playerReceivedAnswers(context))
}

func gameCreated(ctx *Context) func(g cah.Game) {
	return func(g cah.Game) {
		response := &Notification{
			Type: "game_created_event",
			Payload: struct{ Id string }{
				Id: g.Id,
			},
		}

		json, _ := json.Marshal(response)
		ctx.Broadcast <- []byte(json)
	}
}

func roundStarted(ctx *Context) func(payload cah.RoundStartedEventPayload) {
	return func(payload cah.RoundStartedEventPayload) {
		response := &Notification{
			Type:   "round_started_event",
			GameId: payload.Game.Id,
			Payload: struct {
				Round      cah.Round
				FirstRound bool
			}{
				Round:      payload.Round,
				FirstRound: true,
			},
		}

		json, _ := json.Marshal(response)
		ctx.Broadcast <- []byte(json)
	}

}

func playerJoined(ctx *Context) func(payload cah.PlayerJoinedEventPayload) {
	return func(payload cah.PlayerJoinedEventPayload) {
		response := &Notification{
			Type:   "player_joined_event",
			GameId: payload.Game.Id,
			Payload: struct {
				PlayerId   string
				PlayerName string
			}{
				PlayerId:   payload.Player.Id,
				PlayerName: payload.Player.Name,
			},
		}

		json, _ := json.Marshal(response)
		ctx.Broadcast <- []byte(json)
	}
}

func answerPlayed(ctx *Context) func(playerAnswer cah.PlayerAnswer, g cah.Game) {
	return func(playerAnswer cah.PlayerAnswer, g cah.Game) {
		response := &Notification{
			Type:   "answer_played_event",
			Scope:  ONLY_GAME_MEMBERS,
			GameId: g.Id,
			Payload: struct{ PlayerId string }{
				PlayerId: playerAnswer.PlayerId,
			},
		}
		json, _ := json.Marshal(response)
		ctx.Broadcast <- []byte(json)
	}

}

func answerPicked(ctx *Context) func(playerAnswer cah.PlayerAnswer, g cah.Game) {
	return func(playerAnswer cah.PlayerAnswer, g cah.Game) {
		response := &Notification{
			Type:   "answer_picked_event",
			Scope:  ONLY_GAME_MEMBERS,
			GameId: g.Id,
			Payload: struct {
				PlayerId string
				Answers  []cah.Answer
			}{
				PlayerId: playerAnswer.PlayerId,
				Answers:  playerAnswer.Answers,
			},
		}
		json, _ := json.Marshal(response)
		ctx.Broadcast <- []byte(json)
	}
}

func answersRevealed(ctx *Context) func(playerAnswer []cah.PlayerAnswer, g cah.Game) {
	return func(playerAnswers []cah.PlayerAnswer, g cah.Game) {
		response := &Notification{
			Type:   "answers_revealed_event",
			Scope:  ONLY_GAME_MEMBERS,
			GameId: g.Id,
			Payload: struct{ Answers []cah.PlayerAnswer }{
				Answers: playerAnswers,
			},
		}
		json, _ := json.Marshal(response)
		ctx.Broadcast <- []byte(json)
	}

}

func playerReceivedAnswers(ctx *Context) func(payload cah.PlayerReceivedAnswersEventPayload) {
	return func(payload cah.PlayerReceivedAnswersEventPayload) {
		response := &Notification{
			Type:  "answers_received_event",
			Scope: ALL,
			Payload: struct{ Answers []cah.Answer }{
				Answers: payload.Answers,
			},
		}
		json, _ := json.Marshal(response)
		ctx.Broadcast <- []byte(json)
	}
}
