package ws

import cah "github.com/royalzsoftware/cah/src"

func HandleCommand(request *ClientRequest, player *cah.Player) interface{} {
	allCommands := []Command{
		&StartGameCommand{},
		&CreateGameCommand{},
		&LoginCommand{},
		&JoinGameCommand{},
		&GetCardsCommand{},
	}
	com := request.Command
	params := request.Params
	for _, element := range allCommands {
		if element.Command() == com {
			if element.Command() != "login" && player == nil {
				return NotLoggedIn()
			}
			return element.Execute(params, player)
		}
	}
	return CommandNotFound()
}
