package src

type DeckProvider interface {
	getDeck()
}

type DefaultDeckProvider struct {
	DeckProvider
}
