package core

import "testing"

func TestSimpleSpawningInOrder(t *testing.T) {

	// we assume that all possible number spawns go in order 2 -> 3 -> 5 -> 2 ....

	var shouldBeAfterSpawn [4][4]int
	game := NewGame()

	if game.NextSpawn() != 2 {
		t.Errorf("Wrong scheduled spawn, got %v expected %v", game.NextSpawn(), 1)
	}

	game.SpawnVertical()

	shouldBeAfterSpawn = [4][4]int{
		[4]int{2, 0, 0, 0},
		[4]int{0, 0, 2, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	}

	if game.Positions() != shouldBeAfterSpawn {
		t.Errorf("Wrong spawn - received %v instead of %v.", game.Positions(), shouldBeAfterSpawn)
	}

	if game.NextSpawn() != 3 {
		t.Errorf("Wrong scheduled spawn, got %v expected %v", game.NextSpawn(), 2)
	}

	game.moveNumber += 1
	game.SpawnVertical()

	shouldBeAfterSpawn = [4][4]int{
		[4]int{2, 3, 0, 0},
		[4]int{0, 0, 2, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	}

	if game.Positions() != shouldBeAfterSpawn {
		t.Errorf("Wrong spawn - received %v instead of %v.", game.Positions(), shouldBeAfterSpawn)
	}

	if game.NextSpawn() != 5 {
		t.Errorf("Wrong scheduled spawn, got %v expected %v", game.NextSpawn(), 5)
	}

	game.moveNumber += 1
	game.SpawnVertical()

	shouldBeAfterSpawn = [4][4]int{
		[4]int{2, 3, 5, 0},
		[4]int{0, 0, 2, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	}

	if game.Positions() != shouldBeAfterSpawn {
		t.Errorf("Wrong spawn - received %v instead of %v.", game.Positions(), shouldBeAfterSpawn)
	}

	if game.NextSpawn() != 2 {
		t.Errorf("Wrong scheduled spawn, got %v expected %v", game.NextSpawn(), 1)
	}

	game.moveNumber += 1
	game.SpawnVertical()

	shouldBeAfterSpawn = [4][4]int{
		[4]int{2, 3, 5, 2},
		[4]int{0, 0, 2, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	}

	if game.Positions() != shouldBeAfterSpawn {
		t.Errorf("Wrong spawn - received %v instead of %v.", game.Positions(), shouldBeAfterSpawn)
	}

	if game.NextSpawn() != 3 {
		t.Errorf("Wrong scheduled spawn, got %v expected %v", game.NextSpawn(), 2)
	}

	game.moveNumber += 1
	game.SpawnVertical()

	shouldBeAfterSpawn = [4][4]int{
		[4]int{2, 3, 5, 2},
		[4]int{3, 0, 2, 0},
		[4]int{0, 0, 0, 0},
		[4]int{0, 0, 0, 0},
	}

	if game.Positions() != shouldBeAfterSpawn {
		t.Errorf("Wrong spawn - received %v instead of %v.", game.Positions(), shouldBeAfterSpawn)
	}
}

func TestHardSpawningCase(t *testing.T) {
	var shouldBeAfterSpawn [4][4]int
	game := NewGame()

	// First two columns are occupied - so it should be placed in first OK place
	game.data = [4][4]int{
		[4]int{3, 3, 5, 3},
		[4]int{3, 3, 2, 5},
		[4]int{1, 3, 0, 0},
		[4]int{3, 3, 0, 0},
	}

	shouldBeAfterSpawn = [4][4]int{
		[4]int{3, 3, 5, 3},
		[4]int{3, 3, 2, 5},
		[4]int{1, 3, game.NextSpawn(), 0},
		[4]int{3, 3, 0, 0},
	}

	game.SpawnVertical()

	if game.Positions() != shouldBeAfterSpawn {
		t.Errorf("Wrong spawn - received %v instead of %v.", game.Positions(), shouldBeAfterSpawn)
	}
}

func TestOverflowSpawningCase(t *testing.T) {
	var shouldBeAfterSpawn [4][4]int
	game := NewGame()

	game.SpawnVertical()
	game.SpawnVertical()

	// Next spawn should go to third row, but will go to
	// second, as it will be only one with free space
	game.data = [4][4]int{
		[4]int{3, 3, 5, 3},
		[4]int{3, 0, 2, 0},
		[4]int{1, 3, 2, 1},
		[4]int{3, 3, 2, 1},
	}

	shouldBeAfterSpawn = [4][4]int{
		[4]int{3, 3, 5, 3},
		[4]int{3, game.NextSpawn(), 2, 0},
		[4]int{1, 3, 2, 1},
		[4]int{3, 3, 2, 1},
	}

	game.SpawnVertical()

	if game.Positions() != shouldBeAfterSpawn {
		t.Errorf("Wrong spawn - received %v instead of %v.", game.Positions(), shouldBeAfterSpawn)
	}
}
