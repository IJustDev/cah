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

func TestGameAllPlayersShouldHaveEightCards(t *testing.T) {
	g := SetUpDefaultGame()
	AssertEqual(t, len(g.Players[0].Answers), 8, "Player one has: "+strconv.Itoa(len(g.Players[0].Answers)))
	AssertEqual(t, len(g.Players[1].Answers), 8, "Player two has: "+strconv.Itoa(len(g.Players[0].Answers)))
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

func TestZarShouldNotBeAbleToLayAnswers(t *testing.T) {
	g := SetUpDefaultGame()
	AssertEqual(t, g.CurrentRound.Zar.LayAnswers(
		[]cah.Answer{
			g.CurrentRound.Zar.Answers[0],
		}), false, "Zar should not be able to lay an answer")
}

func TestPlayerShouldBeAbleToLayAnswers(t *testing.T) {
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

}

func TestPlugin(t *testing.T) {
	f := &cah.PluginManager{}

	f.LoadAllPlugins("../")

	SetUpDefaultGame()
}
