package problems

// Searches for a word in adjacent cells inside an MxN grid
type SearchWordInGrid struct {
	Grid [][]byte
}

type coord struct {
	root int
	x, y int
}

type node struct {
	letter byte
	children []node
}

func NewSearchWordInGrid() (*SearchWordInGrid, error) {
	return &SearchWordInGrid{
		Grid: [][]byte{
			{'A', 'B', 'C', 'B'},
			{'S', 'F', 'B', 'S'},
			{'A', 'D', 'E', 'E'},
		},
	}, nil
}

func (instance *SearchWordInGrid) SearchWord(word string) bool {

	searchPaths := make([][]coord, len(word))
	for i := range searchPaths {
		searchPaths[i] = make([]coord, 0)
	}

	var tracker []map[coord]bool
	tracker = make([]map[coord]bool, 0)

	for i := range instance.Grid {
		for j := range instance.Grid[i] {
			if instance.Grid[i][j] == word[0] {
				searchPaths[0] = append(searchPaths[0], coord{y: i, x: j, root: len(searchPaths[0])})
				tracker = append(tracker, map[coord]bool{})
				tracker[len(tracker)-1][coord{y: i, x: j, root: len(searchPaths[0]) - 1}] = true
			}
		}
	}

	currentIndex := 0
	for {
		for _, v := range searchPaths[currentIndex] {

			if v.x < (len(instance.Grid[0]) - 1) {
				if instance.Grid[v.y][v.x+1] == word[currentIndex+1] && !tracker[v.root][coord{x: v.x + 1, y: v.y, root: v.root}] {
					searchPaths[currentIndex+1] = append(searchPaths[currentIndex+1], coord{x: v.x + 1, y: v.y, root: v.root})
					tracker[v.root][coord{x: v.x + 1, y: v.y, root: v.root}] = true
				}
			}

			if v.x >= 1 {
				if instance.Grid[v.y][v.x-1] == word[currentIndex+1] && !tracker[v.root][coord{x: v.x - 1, y: v.y, root: v.root}] {
					searchPaths[currentIndex+1] = append(searchPaths[currentIndex+1], coord{x: v.x - 1, y: v.y, root: v.root})
					tracker[v.root][coord{x: v.x - 1, y: v.y, root: v.root}] = true
				}
			}

			if v.y < (len(instance.Grid) - 1) {
				if instance.Grid[v.y+1][v.x] == word[currentIndex+1] && !tracker[v.root][coord{x: v.x, y: v.y + 1, root: v.root}] {
					searchPaths[currentIndex+1] = append(searchPaths[currentIndex+1], coord{x: v.x, y: v.y + 1, root: v.root})
					tracker[v.root][coord{x: v.x, y: v.y + 1, root: v.root}] = true
				}
			}

			if v.y >= 1 {
				if instance.Grid[v.y-1][v.x] == word[currentIndex+1] && !tracker[v.root][coord{x: v.x, y: v.y - 1, root: v.root}] {
					searchPaths[currentIndex+1] = append(searchPaths[currentIndex+1], coord{x: v.x, y: v.y - 1, root: v.root})
					tracker[v.root][coord{x: v.x, y: v.y - 1, root: v.root}] = true
				}
			}
		}

		currentIndex++

		if currentIndex+1 >= len(word) {
			break
		}
	}

	found := false
	for _, v := range tracker  {


		fow _, l := range word {
			
			letterFound := false
			for k, _ := range v {
				if instance.Grid[k.y][k.x] == l { 
					letterFound = true
					continue
				}
			}

			if !letterFound {

			}
		}



	}

	return found
}
