package problems

// Searches for a word in adjacent cells inside an MxN grid
type SearchWordInGrid struct {
	Grid [][]byte
}

type coord struct {
	root int
	x, y int
}

func NewSearchWordInGrid() (*SearchWordInGrid, error) {
	return &SearchWordInGrid{
		Grid: [][]byte{
			{'L', 'I', 'X', 'L'},
			{'P', 'N', 'K', 'H'},
			{'O', 'K', 'I', 'L'},
			{'I', 'E', 'N', 'F'},
		},
	}, nil
}

func (instance *SearchWordInGrid) SearchWord(word string) bool {

	searchPaths := make([][]coord, len(word))
	for i := range searchPaths {
		searchPaths[i] = make([]coord, 0)
	}

	for i := range instance.Grid {
		for j := range instance.Grid[i] {
			if instance.Grid[i][j] == word[0] {
				searchPaths[0] = append(searchPaths[0], coord{y: i, x: j, root: len(searchPaths[0])})
			}
		}
	}

	tracker := map[int]int{}
	currentIndex := 0
	for {
		for _, v := range searchPaths[currentIndex] {
			found := false

			if v.x < (len(instance.Grid[0]) - 1) {
				if instance.Grid[v.y][v.x+1] == word[currentIndex+1] {
					searchPaths[currentIndex+1] = append(searchPaths[currentIndex+1], coord{x: v.x + 1, y: v.y, root: v.root})
					found = true
				}
			}

			if v.x >= 1 {
				if instance.Grid[v.y][v.x-1] == word[currentIndex+1] {
					searchPaths[currentIndex+1] = append(searchPaths[currentIndex+1], coord{x: v.x - 1, y: v.y, root: v.root})
					found = true
				}
			}

			if v.y < (len(instance.Grid) - 1) {
				if instance.Grid[v.y+1][v.x] == word[currentIndex+1] {
					searchPaths[currentIndex+1] = append(searchPaths[currentIndex+1], coord{x: v.x, y: v.y + 1, root: v.root})
					found = true
				}
			}

			if v.y >= 1 {
				if instance.Grid[v.y-1][v.x] == word[currentIndex+1] {
					searchPaths[currentIndex+1] = append(searchPaths[currentIndex+1], coord{x: v.x, y: v.y - 1, root: v.root})
					found = true
				}
			}

			if found {
				tracker[v.root] += 1
			}
		}

		currentIndex++

		if currentIndex+1 >= len(word) {
			break
		}
	}

	for _, s := range tracker {
		if s >= len(word)-1 {
			return true
		}
	}

	return false
}
