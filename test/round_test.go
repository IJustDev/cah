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
		[]cah.Answer{p2.Answers[0]},
	)
	AssertEqual(t, g.CurrentRound.State, 0, "Round state should be 0 -> Laying")
	p3.LayAnswers(
		[]cah.Answer{p3.Answers[0]},
	)
	AssertEqual(t, g.CurrentRound.State, 1, "Round state should be 1 -> Picking")

}
