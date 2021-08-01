package src

type Round struct {
	Question Question
	// All the answers players laid in this specific round
	Answers []PlayerAnswer
	// Will be set after the Zar picked the winning card
	WinningAnswer PlayerAnswer
	// The player who is going to pick the funniest card
	Zar *Player
	// laying = 0, picking = 1, recap = 2
	State int
}

func NewRound(question Question, zar *Player) *Round {
	r := &Round{
		Question: question,
		Zar:      zar,
		Answers:  []PlayerAnswer{},
	}

	return r
}

func (r *Round) AnswerPlayed(playerAnswer PlayerAnswer) {
	r.Answers = append(r.Answers, playerAnswer)
}

func (r *Round) AnswerPicked(winningAnswer PlayerAnswer, player Player) bool {
	if r.State != 1 || r.Zar.Id != player.Id {
		return false
	}
	r.WinningAnswer = winningAnswer
	r.State = 2
	return true
}
func (r *Round) ChangeState(newState int) bool {
	if r.State == newState {
		return false
	}
	r.State = newState
	return true
}
