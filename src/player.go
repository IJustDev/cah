package src

var AllPlayers []*Player

func GetPlayerById(id string) *Player {
	for _, element := range AllPlayers {
		if element.Id == id {
			return element
		}
	}
	return nil
}

type Player struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Answers     []Answer
	CurrentGame *Game
	IsZar       bool
}

func NewPlayer(name string) *Player {
	return &Player{
		Id:   NewObjectId(),
		Name: name,
	}
}

func (p *Player) Join(game *Game) {
	game.Players = append(game.Players, p)
	p.CurrentGame = game
	PlayerJoinedEvent.Trigger(
		PlayerJoinedEventPayload{
			Player: *p,
			Game:   *game,
		},
	)
}

func (p *Player) LayAnswers(answers []Answer) bool {
	if p.IsZar {
		return false
	}
	AnswerPlayedEvent.Trigger(PlayerAnswer{
		PlayerId: p.Id,
		Answers:  answers,
	}, *p.CurrentGame)
	return true
}

func (p *Player) PickCard(playerAnswer PlayerAnswer) bool {
	if !p.IsZar {
		return false
	}
	AnswerPickedEvent.Trigger(playerAnswer, *p.CurrentGame)
	return true
}
