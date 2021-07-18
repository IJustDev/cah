package ws

import cah "github.com/royalzsoftware/cah/src"

type Command interface {
	Command() string
	Execute(map[string]string, *cah.Player) Response
}

type Response struct {
	ErrorCode int
	Details   interface{}
	Data      interface{}
}

func Success(data interface{}) Response {
	return Response{
		ErrorCode: 0,
		Data:      data,
	}
}

func NotLoggedIn() Response {
	return Response{
		ErrorCode: 1,
	}
}

func CommandNotFound() Response {
	return Response{
		ErrorCode: 2,
		Details:   "Command not found",
	}
}

func NotInGame() Response {
	return Response{
		ErrorCode: 3,
		Details:   "Not ingame",
	}
}

func AlreadyInGame() Response {
	return Response{
		ErrorCode: 4,
		Details:   "Already ingame",
	}
}

func NotFound() Response {
	return Response{
		ErrorCode: 5,
		Details:   "Not found",
	}
}
