package ws

import cah "github.com/royalzsoftware/cah/src"

type LoginCommand struct {
}

func (c *LoginCommand) Command() string {
	return "login"
}

func (c *LoginCommand) Execute(params map[string]string, player *cah.Player) interface{} {
	username := params["username"]
	p := cah.NewPlayer(username)
	return NewPlayerCommandResult(*p).Response()
}

type JoinGameCommand struct {
}

func (c *JoinGameCommand) Command() string {
	return "join_game"
}

func (c *JoinGameCommand) Execute(params map[string]string, player *cah.Player) interface{} {
	return params["gameId"]
}
