package core

// import "fmt"

type Game struct {
	data               [GameFieldSize][GameFieldSize]int
	verticalIncoming   int
	horizontalIncoming int
}

type ImpossibleMoveError struct{ error }
type UnknownMoveError struct{ error }

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

func (self *Game) moveRight() (moved bool) {
	for i, row := range self.data {
		newRow := moveRowDown(row[:])
		if newRow != nil {
			moved = true
			for j := 0; j < GameFieldSize; j++ {
				self.data[i][j] = newRow[j]
			}
		}
	}

	return
}

func (self *Game) moveLeft() (moved bool) {
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

	return
}
func (self *Game) moveDown() (moved bool) {
	for i, _ := range self.data {
		thisColumn := make([]int, 4)
		for j := 0; j < GameFieldSize; j++ {
			thisColumn[j] = self.data[j][i]
		}

		newColumn := moveRowDown(thisColumn)

		if newColumn != nil {
			moved = true

			for j := 0; j < GameFieldSize; j++ {
				self.data[j][i] = newColumn[j]
			}
		}
	}

	return
}

func (self *Game) moveUp() (moved bool) {
	for i, _ := range self.data {
		thisColumn := make([]int, 4)
		for j := 0; j < GameFieldSize; j++ {
			thisColumn[j] = self.data[j][i]
		}

		newColumn := moveRowDown(revertSlice(thisColumn))

		if newColumn != nil {
			newColumn = revertSlice(newColumn)
			moved = true

			for j := 0; j < GameFieldSize; j++ {
				self.data[j][i] = newColumn[j]
			}
		}
	}
	return
}

func (self *Game) Move(where int) error {
	var moved bool = false

	switch where {
	case RIGHT:
		moved = self.moveRight()
	case LEFT:
		moved = self.moveLeft()
	case DOWN:
		moved = self.moveDown()
	case UP:
		moved = self.moveUp()
	default:
		return &UnknownMoveError{}
	}

	if moved {
		return nil
	} else {
		return &ImpossibleMoveError{}
	}
}

func (self Game) HaveLost() (bool, error) {
	for _, move := range &[4]int{RIGHT, LEFT, DOWN, UP} {
		moveResult := self.Move(move)

		if moveResult == nil { // move OK
			return false, nil
		}

		switch moveResult.(type) {
		case *ImpossibleMoveError:
			continue
		default:
			return false, moveResult
		}
	}

	return true, nil
}

func (self Game) GetScore() (points int) {
	for _, row := range self.data {
		for _, item := range row {
			if item < 5 {
				continue
			}

			points += item
		}
	}

	return
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
