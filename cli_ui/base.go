package cli_ui

import "fmt"
import "github.com/brain-geek/fivez/core"

func StartGameUI() {
	fmt.Println("Game goes here")

	game := core.NewGame()

	draw(game)

mainloop:
	for {

		loss, _ := game.HaveLost()

		if loss {
			fmt.Println("Thank you for playing!")
			fmt.Printf("Your score is %d.", game.GetScore())
			fmt.Println("")

			break mainloop
		}

		charCode, controlCode, _ := getChar()

		switch controlCode {
		case 0: //KeyEsc:
			if (charCode == 27 /* ESC */) || (charCode == 113 /* q */) || (charCode == 3 /* ctrl-c */) {
				fmt.Println("Thank you for playing!")
				break mainloop
			}
		case 37: //KeyArrowLeft:
			if game.Move(core.LEFT) == nil { // no errors means move was OK
				game.SpawnVertical()
			}
		case 39: //KeyArrowRight:
			if game.Move(core.RIGHT) == nil { // no errors means move was OK
				game.SpawnVertical()
			}
		case 38: //KeyArrowUp:
			if game.Move(core.UP) == nil { // no errors means move was OK
				game.SpawnVertical()
			}
		case 40: //KeyArrowDown:
			if game.Move(core.DOWN) == nil { // no errors means move was OK
				game.SpawnVertical()
			}
		default:
		}

		// Clearing screen
		fmt.Print("\033[H\033[2J")
		draw(game)
	}
}

func draw(game *core.Game) {
	fmt.Println("/===========================\\")

	for i, row := range game.Positions() {
		fmt.Printf("| %4d | %4d | %4d | %4d |", row[0], row[1], row[2], row[3])
		fmt.Println("")

		if i != (core.GameFieldSize - 1) {
			fmt.Println("|===========================|")
		}
	}

	fmt.Println("|===========================|")

	fmt.Println("| Next number: |     ", game.NextSpawn(), "    |")

	fmt.Printf("| Score: |     %9d    |", game.GetScore())
	fmt.Println("")

	fmt.Println("\\===========================/")
}
