package core

import "testing"

func TestScore(t *testing.T) {
	game := NewGame()

	if game.GetScore() != 0 {
		t.Errorf("Expected start game score to be 0, but got %v.", game.GetScore())
	}

	game.data[0][0] = 1

	if game.GetScore() != 0 {
		t.Errorf("Expected start game score to be 0, but got %v.", game.GetScore())
	}

	game.data[0][1] = 3

	if game.GetScore() != 0 {
		t.Errorf("Expected start game score to be 0, but got %v.", game.GetScore())
	}

	game.data[0][2] = 5

	if game.GetScore() != 5 {
		t.Errorf("Expected game with one five to be 5, but got %v.", game.GetScore())
	}

	game.data[0][3] = 10

	if game.GetScore() != 15 {
		t.Errorf("Expected game to sum all stuff, but got %v.", game.GetScore())
	}
}
