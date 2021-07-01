package src

type PlayerAnswer struct {
	Id       string
	PlayerId string
	Answers  []Answer
}

func NewPlayerAnswer(playerId string, answers []Answer) *PlayerAnswer {
	return &PlayerAnswer{
		PlayerId: playerId,
		Answers:  answers,
	}
}

func (p PlayerAnswer) PrepareWithOutPlayerId() *PlayerAnswer {
	return &PlayerAnswer{
		Id:       p.Id,
		PlayerId: "",
		Answers:  p.Answers,
	}
}
