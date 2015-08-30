package core

func (game Game) NextSpawn() int {
	return game.possibleDroppedFigures()[game.spawnCounter%len(game.possibleDroppedFigures())]
}

func (game *Game) SpawnVertical() bool {
	numberToSpawn := game.NextSpawn()
	columnToUseWhenSpawning := game.moveNumber % GameFieldSize

	for i := 0; i < GameFieldSize; i++ {
		for j := 0; j < GameFieldSize; j++ {
			currentColumn := (columnToUseWhenSpawning + j) % GameFieldSize

			if game.data[i][currentColumn] == 0 {
				game.data[i][currentColumn] = numberToSpawn
				game.spawnCounter += 1
				return true
			}
		}

	}

	return false
}
