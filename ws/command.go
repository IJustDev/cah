package ws

import cah "github.com/royalzsoftware/cah/src"

type Command interface {
	Command() string
	Execute(map[string]string, *cah.Player) interface{}
}

type CommandResult interface {
	Response() string
}

type StringCommandResult struct {
	Value string
}

func NewStringCommandResult(response string) *StringCommandResult {
	return &StringCommandResult{
		Value: response,
	}
}

func (c StringCommandResult) Response() string {
	return c.Value
}

type GameCommandResult struct {
	Value cah.Game
}

func NewGameCommandResult(game cah.Game) *GameCommandResult {
	return &GameCommandResult{
		Value: game,
	}
}

func (c GameCommandResult) Response() string {
	return c.Value.Id
}

type PlayerCommandResult struct {
	Value cah.Player
}

func NewPlayerCommandResult(player cah.Player) *PlayerCommandResult {
	return &PlayerCommandResult{
		Value: player,
	}
}

func (c PlayerCommandResult) Response() string {
	return c.Value.Id
}
