package test

import (
	"testing"

	cah "github.com/royalzsoftware/cah/src"
)

func TestRoundStateShouldSwitchAfterAllPlayersLaidAnswers(t *testing.T) {
	g := SetUpDefaultGame()
	p2 := g.Players[1]
	p3 := g.Players[2]

	p2.LayAnswers(
		[]cah.Answer{g.Players[1].Answers[0]},
	)
	AssertEqual(t, g.CurrentRound.State, 0, "Round state should be 0 -> Laying")
	p3.LayAnswers(
		[]cah.Answer{g.Players[2].Answers[0]},
	)
	AssertEqual(t, g.CurrentRound.State, 1, "Round state should be 1 -> Picking")

}

func TestZarShouldNotBeAbleToPickInLayingRoundState(t *testing.T) {
	cah.AnswersRevealedEvent.Register(func(playerAnswers []cah.PlayerAnswer, g cah.Game) {
		zar := g.Players[0]
		x := zar.PickCard(playerAnswers[0])
		AssertEqual(t, x, false, "Zar should not be able to pick an answer")
	})
	SetUpDefaultGame()

}

func TestPickCardShouldChangeRoundState(t *testing.T) {
	cah.AnswersRevealedEvent.Register(func(playerAnswers []cah.PlayerAnswer, g cah.Game) {
		zar := g.Players[0]
		AssertEqual(t, zar.PickCard(playerAnswers[0]), true, "Zar should be able to pick an answer")

		AssertEqual(t, g.CurrentRound.State, 2, "Round state should be 1 -> Picking")
	})
	SetUpCustomGame(3, 1)

}
