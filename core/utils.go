package core

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
