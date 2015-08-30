package core

func (game Game) NextSpawn() int {
	return game.possibleDroppedFigures()[game.spawnCounter%len(game.possibleDroppedFigures())]
}

func (game *Game) SpawnVertical() bool {
	numberToSpawn := game.NextSpawn()
	columnToUseWhenSpawning := game.moveNumber % GameFieldSize

	for i := 0; i < GameFieldSize; i++ {
		if game.data[i][columnToUseWhenSpawning] == 0 {
			game.data[i][columnToUseWhenSpawning] = numberToSpawn
			game.spawnCounter += 1
			return true
		}
	}

	return false
}
