package test

import (
	"fmt"
	"strconv"
	"testing"

	cah "github.com/royalzsoftware/cah/src"
	"github.com/royalzsoftware/cah/ws"
)

func TestCreateGameCommandShouldSucceedAndPlayerShouldHaveJoined(t *testing.T) {
	params := make(map[string]string)
	p := cah.NewPlayer("Player")
	cmd := ws.CreateGameCommand{}
	result := cmd.Execute(params, p)
	payload := result.Data.(ws.CreateGameCommandPayload)
	g := cah.FindGameById(payload.Id)
	AssertNotNil(t, g, "Game with id "+g.Id+" should not be nil")
	AssertNil(t, g.CurrentRound, fmt.Sprintf("%+v", g.CurrentRound))
	found := false
	for _, player := range g.Players {
		if player == p {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Player should have joined to the game.")
	}
}

func TestCreateGameWhilePlayerIsAlreadyInGameShouldFail(t *testing.T) {
	params := make(map[string]string)
	p := cah.NewPlayer("Player")
	g := SetUpDefaultGame()
	p.Join(g)
	cmd := ws.CreateGameCommand{}

	result := cmd.Execute(params, p)
	AssertEqual(t, result.ErrorCode, 4, "ErrorCode should equal 3")
}

func TestStartGameIfHasJoinedToGameShouldWork(t *testing.T) {
	params := make(map[string]string)
	g := SetUpDefaultGame()
	params["gameId"] = g.Id
	cmd := ws.StartGameCommand{}

	result := cmd.Execute(params, g.Players[0])
	AssertEqual(t, result.ErrorCode, 0, "Expected ErrorCode: 0; Actual: "+strconv.Itoa(result.ErrorCode))
}

func TestStartGameIfPlayerDidntJoinAnyGameGameShouldNotWork(t *testing.T) {
	params := make(map[string]string)
	g := SetUpDefaultGame()
	params["gameId"] = g.Id
	cmd := ws.StartGameCommand{}
	p := cah.NewPlayer("PlayerThatWantsToStartTheGameWithoutBeingAPartOfItLul")

	result := cmd.Execute(params, p)
	AssertEqual(t, result.ErrorCode, 3, "Player outside the game should not be able to start the game")
}

func TestJoinGameIfAlreadyInGame(t *testing.T) {
	params := make(map[string]string)
	g := SetUpDefaultGame()
	params["gameId"] = g.Id
	cmd := ws.JoinGameCommand{}
	p := cah.NewPlayer("player")
	AssertNil(t, p.CurrentGame, "The current game should be nil")
	AssertEqual(t, cmd.Execute(params, p).ErrorCode, 0, "Command should have succeeded")
	errorCode := cmd.Execute(params, p).ErrorCode
	AssertEqual(t, errorCode, 4, "Expected: 4; Actual: "+strconv.Itoa(errorCode))
}

func TestGetCardsCommandShouldSucceed(t *testing.T) {
	params := make(map[string]string)
	g := SetUpDefaultGame()
	p := g.Players[0]
	cmd := ws.GetCardsCommand{}
	response := cmd.Execute(params, p).Data.(ws.GetCardsCommandPayload)

	expectedAmountZar := cah.START_ANSWERS - g.CurrentRound.Question.PlaceholderAmount

	AssertEqual(t,
		len(response.Answers),
		expectedAmountZar,
		"Zar should have "+
			strconv.Itoa(
				expectedAmountZar,
			)+
			" answers",
	)
}

func TestPlayCardWhileNotInGame(t *testing.T) {
	params := make(map[string]string)
	params["cardId"] = "invalidId"
	cmd := ws.PlayCardGameCommand{}
	g := SetUpDefaultGame()
	p := g.Players[1] // Is Player
	res := cmd.Execute(params, p)

}
