package ws

import cah "github.com/royalzsoftware/cah/src"

type LoginCommand struct {
}

func (c *LoginCommand) Command() string {
	return "login"
}

func (c *LoginCommand) Execute(params map[string]string, player *cah.Player) Response {
	if _, ok := params["username"]; !ok {
		return InvalidRequest()
	}
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
	if _, ok := params["gameId"]; !ok {
		return InvalidRequest()
	}
	if player.CurrentGame == nil {
		game := cah.FindGameById(params["gameId"])
		player.Join(game)
		return Success(nil)
	}
	return AlreadyInGame()
}

type PlayCardCommand struct {
}

func (c *PlayCardCommand) Command() string {
	return "play_card"
}

func (c *PlayCardCommand) Execute(params map[string]string, player *cah.Player) Response {
	if _, ok := params["answerIds"]; !ok {
		return InvalidRequest()
	}

	if player.CurrentGame != nil {
		return Success(nil)
	}
	return NotInGame()
}
