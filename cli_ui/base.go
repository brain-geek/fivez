package cli_ui

import "fmt"
import "experiments/fivez/core"

func StartGameUI() {
	fmt.Println("Game goes here")

	game := core.NewGame()

	draw(game)

mainloop:
	for {
		charCode, controlCode, _ := getChar()

		switch controlCode {
		case 0: //KeyEsc:
			if (charCode == 27 /* ESC */) || (charCode == 113 /* q */) || (charCode == 3 /* ctrl-c */) {
				fmt.Println("Thank you for playing!")
				break mainloop
			}
		case 37: //KeyArrowLeft:
			game.Move(core.LEFT)
			game.SpawnVertical()
		case 39: //KeyArrowRight:
			game.Move(core.RIGHT)
			game.SpawnVertical()
		case 38: //KeyArrowUp:
			game.Move(core.UP)
			game.SpawnVertical()
		case 40: //KeyArrowDown:
			game.Move(core.DOWN)
			game.SpawnVertical()
		default:
		}
		draw(game)
	}
}

func draw(game *core.Game) {
	fmt.Print("\033[H\033[2J")

	fmt.Println("/===========================\\")

	for i, row := range game.Positions() {
		fmt.Printf("| %4d | %4d | %4d | %4d | \n\r", row[0], row[1], row[2], row[3])

		if i != (core.GameFieldSize - 1) {
			fmt.Println("|===========================|")
		}
	}

	fmt.Println("|===========================|")

	fmt.Println("| Next number: |     ", game.NextSpawn(), "    |")

	fmt.Printf("| Score: |     %9d    |\n\r", game.GetScore())

	fmt.Println("\\===========================/")
}
