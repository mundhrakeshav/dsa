package graph

type Element struct {
	X, Y int
}

func RottenOranges(matrix [][]int) int {
	allRotten := make([]Element, 0)
	totalFreshInitial := 0
	for idxI, arr := range matrix {
		for idxJ, element := range arr {
			if element == 2 {
				allRotten = append(allRotten, Element{
					X: idxI,
					Y: idxJ,
				})
			}
			if element == 1 {
				totalFreshInitial++
			}
		}
	}
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	minutes := 0
	freshEncounters := 0
	for len(allRotten) > 0 {
		nextRottens := make([]Element, 0)
		for _, v := range allRotten {
			for i := 0; i < 4; i++ {
				nx := v.X + dx[i]
				ny := v.Y + dy[i]

				if nx < 0 || nx > len(matrix)-1 || ny < 0 || ny > len(matrix[0])-1 {
					continue
				}
				switch matrix[nx][ny] {
				case 1:
					{
						matrix[v.X+dx[i]][v.Y+dy[i]] = 2
						nextRottens = append(nextRottens, Element{X: nx, Y: ny})
						freshEncounters++
					}
				default:
				}
			}
		}
		if len(nextRottens) > 0 {
			minutes++
		}
		allRotten = nextRottens
	}
	if totalFreshInitial == freshEncounters {
		return minutes
	}
	return -1
}