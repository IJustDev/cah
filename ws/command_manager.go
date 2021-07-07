package ws

import "strings"

func HandleCommand(command string) string {
	allCommands := []Command{
		&StartGameCommand{},
		&LoginCommand{},
	}
	s := strings.Split(command, " ")
	com := s[0]
	params := s[1:]
	for _, element := range allCommands {
		if element.Command() == com {
			if len(element.Parameters()) == len(params) {
				paramsMap := generateParamsMap(element, params)
				return element.Execute(paramsMap)
			}
		}
	}
	return "1"
}

func generateParamsMap(command Command, params []string) map[string]string {
	paramsMap := map[string]string{}
	index := 0
	for _, el := range params {
		paramsMap[el] = params[index]
		index = index + 1
	}
	return paramsMap
}
