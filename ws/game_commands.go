package ws

import cah "github.com/royalzsoftware/cah/src"

type CreateGameCommand struct {
}

func (c *CreateGameCommand) Command() string {
	return "start_game"
}

type CreateGameCommandPayload struct {
	Id string
}

func (c *CreateGameCommand) Execute(params map[string]string, player *cah.Player) Response {
	d := cah.GetDefaultDeck()
	g := cah.NewGame(*d)
	if player.CurrentGame != nil {
		return AlreadyInGame()
	}
	player.Join(g)
	return Success(CreateGameCommandPayload{
		Id: g.Id,
	})
}

// Command use to start an already created game
type StartGameCommand struct {
}

func (c *StartGameCommand) Command() string {
	return "start_game"
}

func (c *StartGameCommand) Execute(params map[string]string, player *cah.Player) Response {
	gameId := params["gameId"]
	game := cah.FindGameById(gameId)
	if game == nil {
		return NotFound()
	}
	if player.CurrentGame == nil || player.CurrentGame.Id != gameId {
		return NotInGame()
	}
	return Success(nil)
}

type GetCardsCommandPayload struct {
	Answers []cah.Answer
}

// Returns all the cards a player has currently has on his hand for his current
// game
type GetCardsCommand struct {
}

func (c *GetCardsCommand) Command() string {
	return "get_cards"
}

func (c *GetCardsCommand) Execute(params map[string]string, player *cah.Player) Response {
	game := player.CurrentGame
	if game == nil {
		return NotInGame()
	}
	return Success(GetCardsCommandPayload{
		Answers: player.Answers,
	})
}
