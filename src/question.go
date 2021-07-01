package src

type Question struct {
	Question          string
	PlaceholderAmount int
}

func (q Question) Fill(answers []Answer) string {
	return q.Question
}
