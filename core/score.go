package core

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
