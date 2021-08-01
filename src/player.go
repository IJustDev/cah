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
}

func NewPlayer(name string) *Player {
	p := &Player{
		Id:   NewObjectId(),
		Name: name,
	}
	AllPlayers = append(AllPlayers, p)
	return p
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
	if p.IsZar() {
		return false
	}
	p.CurrentGame.PlayerLaidAnswers(p, answers)
	AnswerPlayedEvent.Trigger(PlayerAnswer{
		PlayerId: p.Id,
		Answers:  answers,
	}, *p.CurrentGame)
	return true
}

func (p *Player) PickCard(playerAnswer PlayerAnswer) bool {
	if !p.IsZar() {
		return false
	}

	if ok := p.CurrentGame.CurrentRound.AnswerPicked(playerAnswer, *p); ok {
		AnswerPickedEvent.Trigger(playerAnswer, *p.CurrentGame)
		return true
	}
	return false
}

func (p Player) IsZar() bool {
	return p.CurrentGame.CurrentRound.Zar.Name == p.Name
}
