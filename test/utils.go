package test

import (
	"reflect"
	"strconv"
	"testing"

	cah "github.com/royalzsoftware/cah/src"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	t.Fatal(message)
}

func AssertNil(t *testing.T, a interface{}, message string) {
	if reflect.ValueOf(a).IsZero() {
		return
	}
	t.Fatal(message)
}

func AssertNotNil(t *testing.T, a interface{}, message string) {
	if a != nil {
		return
	}
	t.Fatal(message)
}

func SetUpDefaultGame() *cah.Game {
	return SetUpCustomGame(3, 0)
}

func SetUpCustomGame(playerAmount int, state int) *cah.Game {
	deck := cah.GetDefaultDeck()
	g := cah.NewGame(*deck)

	for i := 0; i != playerAmount; i++ {
		p := cah.NewPlayer("P" + strconv.Itoa(i+1))
		p.Join(g)
	}

	g.StartGame()

	if state < 1 {
		return g
	}

	p1 := g.Players[1]
	p2 := g.Players[2]

	p2.LayAnswers([]cah.Answer{
		p2.Answers[0],
	})
	p1.LayAnswers([]cah.Answer{
		p1.Answers[1],
	})

	if state < 2 {
		return g
	}

	return g
}
