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
