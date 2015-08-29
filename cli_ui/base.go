package cli_ui

import "fmt"
import "experiments/fivez/core"

func StartGameUI() {
	fmt.Println("Game goes here")

	game := core.NewGame()

	draw(game)
}

func draw(game *core.Game) {
	fmt.Println("/=======================\\")

	for i, row := range game.Positions() {
		fmt.Println("| ", row[0], " | ", row[1], " | ", row[2], " | ", row[3], " |")

		if i != (core.GameFieldSize - 1) {
			fmt.Println("|=======================|")
		}
	}

	fmt.Println("\\=======================/")
}
