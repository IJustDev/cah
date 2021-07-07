package ws

import cah "github.com/royalzsoftware/cah/src"

type LoginCommand struct {
}

func (c *LoginCommand) Command() string {
	return "login"
}

func (c *LoginCommand) Parameters() []string {
	return []string{"username"}
}

func (c *LoginCommand) Execute(params map[string]string) string {
	username := params["username"]
	p := cah.NewPlayer(username)
	return NewPlayerCommandResult(*p).Response()
}

type JoinGameCommand struct {
}

func (c *JoinGameCommand) Command() string {
	return "start_game"
}

func (c *JoinGameCommand) Parameters() []string {
	return []string{"gameId"}
}

func (c *JoinGameCommand) Execute(params map[string]string) string {
	return ""
}
