package test

import (
	"strconv"
	"testing"

	cah "github.com/royalzsoftware/cah/src"
)

func TestGameShouldContainThreePlayers(t *testing.T) {
	g := SetUpDefaultGame()
	AssertEqual(t, len(g.Players), 3, "Players should be two")
}

func TestGameShouldNotStartWithLessThanThreePlayers(t *testing.T) {
	deck := cah.GetDefaultDeck()
	g := cah.NewGame(*deck)

	p1 := cah.NewPlayer("P1")
	p2 := cah.NewPlayer("P2")
	p1.Join(g)
	p2.Join(g)

	AssertEqual(t, g.StartGame(), false, "Game should not be startable with less than three players")
}

func TestGameStartRoundShouldNotBeNil(t *testing.T) {
	g := SetUpDefaultGame()
	AssertNotNil(t, g.CurrentRound, "Current round should not be nil")
}

func TestGameAllPlayersExceptZarShouldSameAmountOfAnswers(t *testing.T) {
	g := SetUpDefaultGame()

	playerAmount := cah.START_ANSWERS
	placeHolderAmount := g.CurrentRound.Question.PlaceholderAmount
	zarAmount := playerAmount - placeHolderAmount

	AssertEqual(t, len(g.Players[0].Answers), zarAmount, "Player one has: "+strconv.Itoa(len(g.Players[0].Answers)))
	AssertEqual(t, len(g.Players[1].Answers), playerAmount, "Player two has: "+strconv.Itoa(len(g.Players[0].Answers)))
}

func TestGameZarPickOrder(t *testing.T) {
	g := SetUpDefaultGame()
	AssertEqual(t, g.CurrentRound.Zar.Name, "P1", "Player one should have been the zar; "+g.CurrentRound.Zar.Name)
	g.NextRound()
	AssertEqual(t, g.CurrentRound.Zar.Name, "P2", "Player two should have been the zar; "+g.CurrentRound.Zar.Name)
	g.NextRound()
	AssertEqual(t, g.CurrentRound.Zar.Name, "P3", "Player one should have been the zar; "+g.CurrentRound.Zar.Name)
	g.NextRound()
	AssertEqual(t, g.CurrentRound.Zar.Name, "P1", "Player one should have been the zar; "+g.CurrentRound.Zar.Name)
}

func TestPlayerShouldBeAbleToLayCardButZarNot(t *testing.T) {
	g := SetUpDefaultGame()
	AssertEqual(t, g.Players[0].LayAnswers(
		[]cah.Answer{
			g.CurrentRound.Zar.Answers[0],
		}), false, "Zar should not be able to lay an answer")
	AssertEqual(t, g.Players[1].LayAnswers(
		[]cah.Answer{
			g.Players[1].Answers[0],
		}), true, "Player should be able to lay an answer")
	AssertEqual(t, g.Players[2].LayAnswers(
		[]cah.Answer{
			g.Players[2].Answers[0],
		}), true, "Player should be able to lay an answer")

	AssertEqual(t, g.CurrentRound.State, 1, "Round state should have switched to picking.")
}

func TestPlugin(t *testing.T) {
	f := &cah.PluginManager{}

	f.LoadAllPlugins("../")

	SetUpDefaultGame()
}

func TestZarShouldBeAbleToPickACard(t *testing.T) {
	g := SetUpCustomGame(3, 1)
	zar := g.Players[0]

	AssertEqual(
		t,
		zar.PickCard(g.CurrentRound.Answers[0]),
		true,
		"Zar should have been able to pick a card.",
	)
	AssertEqual(
		t,
		g.CurrentRound.State,
		2,
		"Round state should have switched to recap",
	)
}
