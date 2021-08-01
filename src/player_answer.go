package src

var AllAnswers []PlayerAnswer

func GetPlayerAnswerById(id string) *PlayerAnswer {
	for _, element := range AllAnswers {
		if element.Id == id {
			return &element
		}
	}
	return nil
}

type PlayerAnswer struct {
	Id       string
	PlayerId string
	Answers  []Answer
}

func NewPlayerAnswer(playerId string, answers []Answer) *PlayerAnswer {
	answer := &PlayerAnswer{
		PlayerId: playerId,
		Answers:  answers,
	}

	AllAnswers = append(AllAnswers, *answer)
	return answer
}

func (p PlayerAnswer) PrepareWithOutPlayerId() *PlayerAnswer {
	return &PlayerAnswer{
		Id:       p.Id,
		PlayerId: "",
		Answers:  p.Answers,
	}
}
