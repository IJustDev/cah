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

// ErrorCode 0
func Success(data interface{}) Response {
	return Response{
		ErrorCode: 0,
		Data:      data,
	}
}

// ErrorCode 1
func NotLoggedIn() Response {
	return Response{
		ErrorCode: 1,
	}
}

// ErrorCode 2
func CommandNotFound() Response {
	return Response{
		ErrorCode: 2,
		Details:   "Command not found",
	}
}

// ErrorCode 3
func NotInGame() Response {
	return Response{
		ErrorCode: 3,
		Details:   "Not ingame",
	}
}

// ErrorCode 4
func AlreadyInGame() Response {
	return Response{
		ErrorCode: 4,
		Details:   "Already ingame",
	}
}

// ErrorCode 5
func NotFound() Response {
	return Response{
		ErrorCode: 5,
		Details:   "Not found",
	}
}

// ErrorCode 6
func InvalidRequest() Response {
	return Response{
		ErrorCode: 6,
		Details:   "Invalid request",
	}
}
