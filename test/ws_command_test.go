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
