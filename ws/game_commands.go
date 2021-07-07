package ws

import cah "github.com/royalzsoftware/cah/src"

type StartGameCommand struct {
}

func (c *StartGameCommand) Command() string {
	return "start_game"
}

func (c *StartGameCommand) Parameters() []string {
	return []string{}
}

func (c *StartGameCommand) Execute(params map[string]string) string {
	d := cah.GetDefaultDeck()
	g := cah.NewGame(*d)
	return NewGameCommandResult(*g).Response()
}
