package src

var AllGames []*Game

func GetGameById(id string) *Game {
	for _, element := range AllGames {
		if element.Id == id {
			return element
		}
	}
	return nil
}

type Game struct {
	Id           string
	LastRound    *Round
	CurrentRound *Round
	Players      []*Player
	Decks        []Deck
	zarIndex     int
	roundCount   int
	PlayersLaid  int
}

func NewGame(deck Deck) *Game {
	g := &Game{
		Id:       NewObjectId(),
		Decks:    []Deck{deck},
		zarIndex: -1,
	}
	AllGames = append(AllGames, g)
	return g
}

func (g *Game) StartGame() {
	g.CurrentRound = NewRound(g.determineRandomQuestion(), g.determineNewZar())
	g.GivePlayerCards(true, 8)
	GameStartedEvent.Trigger(*g)
}

func (g *Game) GivePlayerCards(giveZar bool, amount int) {
	for _, player := range g.Players {
		for i := 0; i != amount; i++ {
			player.Answers = append(player.Answers, g.getRandomAnswer())
			PlayerReceivedAnswersEvent.Trigger(PlayerReceivedAnswersEventPayload{
				Player:  player,
				Answers: player.Answers,
			})
		}
	}
}

func (g *Game) NextRound() {
	g.LastRound = g.CurrentRound
	g.CurrentRound = NewRound(g.determineRandomQuestion(), g.determineNewZar())
	g.GivePlayerCards(false, g.LastRound.Question.PlaceholderAmount)
	RoundStartedEvent.Trigger(RoundStartedEventPayload{
		Round:      *g.CurrentRound,
		RoundCount: g.roundCount,
		Game:       *g,
	})
}

func (g *Game) getRandomAnswer() Answer {
	return g.Decks[0].Answers[0]
}

func (g *Game) determineNewZar() *Player {
	if g.zarIndex == len(g.Players)-1 {
		g.zarIndex = 0
	} else {
		g.zarIndex++
	}
	return g.Players[g.zarIndex]
}

func (g *Game) determineRandomQuestion() Question {
	return g.Decks[0].Questions[0]
}

func (g *Game) PlayerLaidAnswers(player *Player, answers []Answer) bool {
	g.CurrentRound.Answers = append(g.CurrentRound.Answers,
		*NewPlayerAnswer(
			player.Id,
			answers,
		),
	)
	if len(g.CurrentRound.Answers) == len(g.Players)-1 {
		AnswersRevealedEvent.Trigger(g.CurrentRound.Answers, *g)
		g.CurrentRound.ChangeState(1)
	}
	return true
}
