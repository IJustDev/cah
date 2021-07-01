package src

type Deck struct {
	Questions []Question `json:"questions"`
	Answers   []Answer   `json:"answers"`
}

func GetDefaultDeck() *Deck {
	return &Deck{
		Questions: []Question{
			Question{
				Question: "I like big **.",
			},
		},
		Answers: []Answer{
			Answer{
				Answer: "Butts and I can not lie.",
			},
		},
	}
}
