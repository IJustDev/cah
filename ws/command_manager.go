package ws

import cah "github.com/royalzsoftware/cah/src"

type ErrorResponse struct {
	Error int
}

func HandleCommand(request *ClientRequest, player *cah.Player) interface{} {
	allCommands := []Command{
		&StartGameCommand{},
		&LoginCommand{},
		&JoinGameCommand{},
		&GetCardsCommand{},
	}
	com := request.Command
	params := request.Params
	for _, element := range allCommands {
		if element.Command() == com {
			if element.Command() != "login" && player == nil {
				return &ErrorResponse{Error: 2}
			}
			return element.Execute(params, player)
		}
	}
	return &ErrorResponse{Error: 1}
}
