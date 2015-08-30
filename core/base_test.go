package core

import "testing"
import "reflect"

// import "fmt"

func TestNewGamePositions(t *testing.T) {
	t.Log("Creating new Game")
	game := NewGame()

	for x, i := range game.Positions() {
		for y, value := range i {
			if (x == FirstBlockX) && (y == FirstBlockY) {
				continue
			}

			if value != 0 {
				t.Errorf("Expected score of 0, but it was %d instead.", value)
			}
		}
	}
}

func TestNewGameIncoming(t *testing.T) {
	t.Log("Creating new Game")
	game := NewGame()

	value := game.getVerticalIncoming()
	if value != 0 {
		t.Errorf("Expected vertical incoming of 0, but it was %d instead.", value)
	}

	value = game.getHorizontalIncoming()
	if value != 0 {
		t.Errorf("Expected horizontal incoming of 0, but it was %d instead.", value)
	}
}

func TestPossibleIncoming(t *testing.T) {
	t.Log("Creating new Game")
	game := NewGame()

	possibilities := game.possibleDroppedFigures()
	if !reflect.DeepEqual(possibilities, []int{2, 3, 5}) {
		t.Errorf("Expected possibilities to be 2,3,5. But got something else.")
	}
}

func TestCanMerge(t *testing.T) {
	t.Log("Testing non-mergeable scenarios")

	if canMerge(0, 0) {
		t.Errorf("Expected 0 and 0 to be not mergeable, but was mergeable instead")
	}

	if canMerge(1, 1) {
		t.Errorf("Expected 1 and 1 to be not mergeable, but was mergeable instead")
	}

	if canMerge(2, 2) {
		t.Errorf("Expected 2 and 2 to be not mergeable, but was mergeable instead")
	}

	if canMerge(3, 3) {
		t.Errorf("Expected 3 and 3 to be not mergeable, but was mergeable instead")
	}

	if canMerge(4, 4) {
		t.Errorf("Expected 4 and 4 to be not mergeable, but was mergeable instead")
	}

	if canMerge(2, 5) {
		t.Errorf("Expected 2 and 5 to be not mergeable, but was mergeable instead")
	}

	if canMerge(5, 2) {
		t.Errorf("Expected 5 and 2 to be not mergeable, but was mergeable instead")
	}

	t.Log("Testing mergeable scenarios")

	if !canMerge(2, 3) {
		t.Errorf("Expected 2 and 3 to be mergeable, but was not.")
	}

	if !canMerge(3, 2) {
		t.Errorf("Expected 3 and 2 to be mergeable, but was not.")
	}

	if !canMerge(5, 5) {
		t.Errorf("Expected 5 and 5 to be mergeable, but was not.")
	}

	if !canMerge(100, 100) {
		t.Errorf("Expected 100 and 100 to be mergeable, but was not.")
	}
}

func TestLoss(t *testing.T) {
	game := NewGame()

	result, err := game.HaveLost()
	if result {
		t.Errorf("Expected start position to be non-lossy, but it is.")
	}

	if err != nil {
		t.Errorf("Expected to have no errors, but got %v.", err)
	}

	if game.Positions() != NewGame().Positions() {
		t.Errorf("Expected checking for loss to not change positions, but got %v.", game.Positions())
	}

	game.data = [4][4]int{
		[4]int{1, 5, 2, 5},
		[4]int{100, 10, 100, 10},
		[4]int{200, 20, 200, 20},
		[4]int{1, 30, 2, 30},
	}

	result, err = game.HaveLost()

	if !result {
		t.Errorf("Expected lossy position to be lossy, but it is not.")
	}

	if err != nil {
		t.Errorf("Expected to have no errors, but got %v.", err)
	}
}

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
