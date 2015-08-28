package base

// import "fmt"

type Game struct {
	data               [GameFieldSize][GameFieldSize]int
	verticalIncoming   int
	horizontalIncoming int
}

func NewGame() *Game {
	game := new(Game)
	game.data[FirstBlockX][FirstBlockY] = FirstBlockData

	return game
}

func (self *Game) Positions() [GameFieldSize][GameFieldSize]int {
	return self.data
}

func (self *Game) possibleDroppedFigures() []int {
	return []int{2, 3, 5}
}

func (self *Game) getVerticalIncoming() int {
	return self.verticalIncoming
}

func (self *Game) getHorizontalIncoming() int {
	return self.horizontalIncoming
}

func (self *Game) Move(where int) error {
	var moved bool = false

	switch where {
	case RIGHT:
		for i, row := range self.data {
			newRow := moveRowDown(row[:])
			if newRow != nil {
				moved = true
				for j := 0; j < GameFieldSize; j++ {
					self.data[i][j] = newRow[j]
				}
			}
		}
	case LEFT:
		for i, row := range self.data {
			reverseRow := revertSlice(row[:])
			newReversedRow := moveRowDown(reverseRow)

			if newReversedRow != nil {
				newRow := revertSlice(newReversedRow)

				moved = true

				for j := 0; j < GameFieldSize; j++ {
					self.data[i][j] = newRow[j]
				}
			}
		}
	}

	_ = moved

	return nil
}

func canMerge(first, second int) bool {
	if (first == 0) || (second == 0) {
		return false
	}

	if first+second == FirstHardNumber {
		return true
	}

	if (first >= FirstHardNumber) && (second >= FirstHardNumber) {
		return first == second
	}

	return false
}

func moveRowDown(row []int) (result []int) {
	var moved bool = false

	result = make([]int, len(row))
	copy(result, row)

	for i := GameFieldSize - 1; i > 0; i-- {
		if ((result[i-1] != 0) && (result[i] == 0)) || canMerge(result[i], result[i-1]) {
			moved = true
			result[i] += result[i-1]
			result[i-1] = 0
		}
	}

	if moved {
		return
	} else {
		return nil
	}
}

func revertSlice(original []int) (newSlice []int) {
	oriLen := len(original)
	newSlice = make([]int, oriLen)

	for i := 0; i < oriLen; i++ {
		newSlice[i] = original[oriLen-i-1]
	}

	return
}

const (
	// Game field is 4x4
	GameFieldSize = 4

	// Where to start the first block
	FirstBlockX = 1
	FirstBlockY = 2

	// What will be in the first block
	FirstBlockData = 2

	// Starting from what number we allow only self-merges
	FirstHardNumber = 5

	UP    = 0
	LEFT  = 1
	DOWN  = 2
	RIGHT = 4
)
