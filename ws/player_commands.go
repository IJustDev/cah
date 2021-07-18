package ws

import cah "github.com/royalzsoftware/cah/src"

type LoginCommand struct {
}

func (c *LoginCommand) Command() string {
	return "login"
}

func (c *LoginCommand) Execute(params map[string]string, player *cah.Player) Response {
	username := params["username"]
	p := cah.NewPlayer(username)
	return Success(p.Id)
}

type JoinGameCommand struct {
}

func (c *JoinGameCommand) Command() string {
	return "join_game"
}

func (c *JoinGameCommand) Execute(params map[string]string, player *cah.Player) Response {
	return Success(nil)
}
