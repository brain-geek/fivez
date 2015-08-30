package core

import "testing"
import "reflect"

func TestFailMove(t *testing.T) {
	game := NewGame()

	game.data[1][3] = 2

	e := game.Move(RIGHT)

	if (e == nil) || (reflect.TypeOf(e).String() != "*core.ImpossibleMoveError") {
		t.Errorf("Expected move create ImpossibleMoveError, got %v instead", e)
	}
}

func TestUnknownMove(t *testing.T) {
	game := NewGame()

	game.data[1][3] = 2

	e := game.Move(-123)

	if (e == nil) || (reflect.TypeOf(e).String() != "*core.UnknownMoveError") {
		t.Errorf("Expected move create UnknownMoveError, got %v instead", e)
	}
}

func TestSimpleMoveRight(t *testing.T) {
	game := NewGame()

	e := game.Move(RIGHT)

	if e != nil {
		t.Errorf("Expected move to be OK, got %v instead", e)
	}

	shouldBeArray := [4][4]int{
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 2},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	}

	value := game.Positions() == shouldBeArray

	if !value {
		t.Errorf("Expected existing point to move one row right from %v, but got %v instead.", NewGame().Positions(), game.Positions())
	}

}

func TestSimpleMoveLeft(t *testing.T) {
	game := NewGame()

	e := game.Move(LEFT)

	if e != nil {
		t.Errorf("Expected move to be OK, got %v instead", e)
	}

	shouldBeArray := [4][4]int{
		[4]int{0, 0, 0, 0},
		[4]int{0, 2, 0, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	}

	value := game.Positions() == shouldBeArray

	if !value {
		t.Errorf("Expected existing point to move one row left from %v, but got %v instead.", NewGame().Positions(), game.Positions())
	}

}

func TestSimpleMoveDown(t *testing.T) {
	game := NewGame()

	e := game.Move(DOWN)

	if e != nil {
		t.Errorf("Expected move to be OK, got %v instead", e)
	}

	shouldBeArray := [4][4]int{
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 2, 0},
		[4]int{0, 0, 0, 0},
	}

	value := game.Positions() == shouldBeArray

	if !value {
		t.Errorf("Expected existing point to move one row down from %v, but got %v instead.", NewGame().Positions(), game.Positions())
	}

}

func TestSimpleMoveUp(t *testing.T) {
	game := NewGame()

	e := game.Move(UP)

	if e != nil {
		t.Errorf("Expected move to be OK, got %v instead", e)
	}

	shouldBeArray := [4][4]int{
		[4]int{0, 0, 2, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	}

	value := game.Positions() == shouldBeArray

	if !value {
		t.Errorf("Expected existing point to move one row top from %v, but got %v instead.", NewGame().Positions(), game.Positions())
	}

}

func TestMoveNumbering(t *testing.T) {
	game := NewGame()

	if game.GetMoveNumber() != 0 {
		t.Errorf("Wrong move number, got %v, expected %v.", game.GetMoveNumber(), 0)
	}

	_ = game.Move(DOWN)

	if game.GetMoveNumber() != 1 {
		t.Errorf("Wrong move number, got %v, expected %v.", game.GetMoveNumber(), 1)
	}

	_ = game.Move(DOWN)

	if game.GetMoveNumber() != 2 {
		t.Errorf("Wrong move number, got %v, expected %v.", game.GetMoveNumber(), 2)
	}
}
