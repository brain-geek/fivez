package cli_ui

import "fmt"
import "experiments/fivez/core"

func StartGameUI() {
	fmt.Println("Game goes here")

	game := core.NewGame()

	draw(game)

mainloop:
	for {
		_, char, _ := getChar()

		fmt.Println(char)

		switch char {
		case 0: //KeyEsc:
			break mainloop
		case 37: //KeyArrowLeft:
			game.Move(core.LEFT)
		case 39: //KeyArrowRight:
			game.Move(core.RIGHT)
		case 38: //KeyArrowUp:
			game.Move(core.UP)
		case 40: //KeyArrowDown:
			game.Move(core.DOWN)
		default:
		}
		draw(game)
	}
}

func draw(game *core.Game) {
	fmt.Print("\033[H\033[2J")

	fmt.Println("/=======================\\")

	for i, row := range game.Positions() {
		fmt.Println("| ", row[0], " | ", row[1], " | ", row[2], " | ", row[3], " |")

		if i != (core.GameFieldSize - 1) {
			fmt.Println("|=======================|")
		}
	}

	fmt.Println("\\=======================/")
}
