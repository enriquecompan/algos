package problems

// Searches for a word in adjacent cells inside an MxN grid
type SearchWordInGrid struct {
	Grid [][]byte
}

type node struct {
	letter   byte
	x, y     int
	parent   *node
	children []node
}

func NewSearchWordInGrid() (*SearchWordInGrid, error) {
	return &SearchWordInGrid{
		Grid: [][]byte{
			{'C', 'A', 'A'},
			{'A', 'A', 'A'},
			{'B', 'C', 'D'},
		},
	}, nil
}

func findStringInPath(_node node, word string) string {
	path := dfs(&_node, word, 0)
	if path != nil {
		return reconstructWord(path)
	}
	return ""
}

func dfs(n *node, word string, index int) *node {
	if index >= len(word) || n.letter != word[index] {
		return nil
	}
	if index == len(word)-1 {
		return n
	}
	for i := range n.children {
		if result := dfs(&n.children[i], word, index+1); result != nil {
			return result
		}
	}
	return nil
}

func reconstructWord(n *node) string {
	var chars []byte
	for n != nil {
		chars = append([]byte{n.letter}, chars...)
		n = n.parent
	}
	return string(chars)
}

func (instance *SearchWordInGrid) SearchWord(word string) bool {

	// Initialize the root node of the search paths
	searchPaths := make([]node, 0)
	for i := range instance.Grid {
		for j := range instance.Grid[i] {
			if instance.Grid[i][j] == word[0] {
				searchPaths = append(searchPaths, node{letter: word[0], y: i, x: j, parent: nil, children: []node{}})
			}
		}
	}

	gridWidth := len(instance.Grid[0])
	gridHeight := len(instance.Grid)

	for rootIndex := range searchPaths {
		currentNode := &searchPaths[rootIndex]
		currentChildren := []*node{}
		rootNodeProcessed := false

		for letterIndex := 1; letterIndex < len(word); letterIndex++ {

			currentLetter := word[letterIndex]

			// Root node processing
			if !rootNodeProcessed {
				rootNodeProcessed = true

				if currentNode.x < (gridWidth-1) && instance.Grid[currentNode.y][currentNode.x+1] == byte(currentLetter) {

					currentNode.children = append(currentNode.children,
						node{
							letter: byte(currentLetter),
							x:      currentNode.x + 1,
							y:      currentNode.y,
							parent: currentNode,
						})

				}

				if currentNode.x >= 1 && instance.Grid[currentNode.y][currentNode.x-1] == byte(currentLetter) {

					currentNode.children = append(currentNode.children,
						node{
							letter: byte(currentLetter),
							x:      currentNode.x - 1,
							y:      currentNode.y,
							parent: currentNode,
						})

				}

				if currentNode.y < (gridHeight-1) && instance.Grid[currentNode.y+1][currentNode.x] == byte(currentLetter) {

					currentNode.children = append(currentNode.children,
						node{
							letter: byte(currentLetter),
							x:      currentNode.x,
							y:      currentNode.y + 1,
							parent: currentNode,
						})

				}

				if currentNode.y >= 1 && instance.Grid[currentNode.y-1][currentNode.x] == byte(currentLetter) {

					currentNode.children = append(currentNode.children,
						node{
							letter: byte(currentLetter),
							x:      currentNode.x,
							y:      currentNode.y - 1,
							parent: currentNode,
						})

				}

				for i := range currentNode.children {
					currentChildren = append(currentChildren, &currentNode.children[i])
				}

			} else {
				// Children nodes processing

				nextChildren := []*node{}
				for childIndex := range currentChildren {

					currentNode = currentChildren[childIndex]

					if currentNode.x < (gridWidth-1) && instance.Grid[currentNode.y][currentNode.x+1] == byte(currentLetter) {
						if !(currentNode.x+1 == currentNode.parent.x && currentNode.y == currentNode.parent.y) {
							currentNode.children = append(currentNode.children,
								node{
									letter: byte(currentLetter),
									x:      currentNode.x + 1,
									y:      currentNode.y,
									parent: currentNode,
								})
						}
					}

					if currentNode.x >= 1 && instance.Grid[currentNode.y][currentNode.x-1] == byte(currentLetter) {
						if !(currentNode.x-1 == currentNode.parent.x && currentNode.y == currentNode.parent.y) {
							currentNode.children = append(currentNode.children,
								node{
									letter: byte(currentLetter),
									x:      currentNode.x - 1,
									y:      currentNode.y,
									parent: currentNode,
								})
						}
					}

					if currentNode.y < (gridHeight-1) && instance.Grid[currentNode.y+1][currentNode.x] == byte(currentLetter) {
						if !(currentNode.x == currentNode.parent.x && currentNode.y+1 == currentNode.parent.y) {
							currentNode.children = append(currentNode.children,
								node{
									letter: byte(currentLetter),
									x:      currentNode.x,
									y:      currentNode.y + 1,
									parent: currentNode,
								})
						}
					}

					if currentNode.y >= 1 && instance.Grid[currentNode.y-1][currentNode.x] == byte(currentLetter) {
						if !(currentNode.x == currentNode.parent.x && currentNode.y-1 == currentNode.parent.y) {
							currentNode.children = append(currentNode.children,
								node{
									letter: byte(currentLetter),
									x:      currentNode.x,
									y:      currentNode.y - 1,
									parent: currentNode,
								})
						}
					}

					for i := range currentNode.children {
						nextChildren = append(nextChildren, &currentNode.children[i])
					}
				}

				currentChildren = nextChildren
			}

		}
	}

	result := false
	for _, path := range searchPaths {
		str := findStringInPath(path, word)

		if str == word {
			result = true
			break
		}
	}

	return result
}
