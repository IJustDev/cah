package ws

import cah "github.com/royalzsoftware/cah/src"

type StartGameCommand struct {
}

func (c *StartGameCommand) Command() string {
	return "start_game"
}

func (c *StartGameCommand) Execute(params map[string]string, player *cah.Player) interface{} {
	d := cah.GetDefaultDeck()
	g := cah.NewGame(*d)
	g.StartGame()
	return NewGameCommandResult(*g).Response()
}

type GetCardsCommand struct {
}

func (c *GetCardsCommand) Command() string {
	return "get_cards"
}

func (c *GetCardsCommand) Execute(params map[string]string, player *cah.Player) interface{} {
	return player.Id
}
