package ws

func HandleCommand(command string) string {
	switch command {
	case "TestCommand":
		return "Test command executed"
	default:
		return "No command found"
	}
}
