package graph

func GetProvincesCountRecursive(arr [][]int) int {
	visited := make(map[int]struct{}, len(arr))
	count := 0
	for i := range arr {
		if _, found := visited[i]; !found {
			count++
			traverse(arr, visited, i)
		}
	}
	return count
}

func GetProvincesCountIterative(arr [][]int) int {
	visited := make(map[int]struct{}, len(arr))
	count := 0
	for i := range arr {
		if _, found := visited[i]; !found {
			count++

			cities := []int{i}
			for len(cities) > 0 {
				connectedCities := make([]int, 0)
				for _, city := range cities {
					if _, found := visited[city]; !found {
						visited[city] = struct{}{}
						for j := 0; j < len(arr); j++ {
							if arr[city][j] != 0 {
								connectedCities = append(connectedCities, j)
							}
						}
					}
				}
				cities = connectedCities
			}
		}
	}
	return count
}

func traverse(arr [][]int, visited map[int]struct{}, i int) {
	if _, ok := visited[i]; ok {
		return
	}

	visited[i] = struct{}{}

	for j, adjNode := range arr[i] {
		if adjNode == 1 {
			traverse(arr, visited, j)
		}
	}
}
