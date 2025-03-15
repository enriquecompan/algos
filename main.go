package main

import (
	"fmt"
	"slices"

	"github.com/enriquecompan/algos/internal/algos"
	"github.com/enriquecompan/algos/internal/problems"
)

func main() {
	unsorted := []int{3, 1, 6, 7, 9, 4, 2, 10, 5, 8, 5}

	fmt.Println("\nUnsorted array:", unsorted)

	// Bubblesort
	work := slices.Clone(unsorted)
	algos.BubbleSort(work)
	fmt.Println("\nBubblesort result:", work)

	// Quicksort
	work = slices.Clone(unsorted)
	work = algos.QuickSort(work)
	fmt.Println("Quicksort result:", work)

	// Search word in MxN grid in different patterns
	searchInGrid, _ := problems.NewSearchWordInGrid()
	fmt.Println("\nMxN grid:", searchInGrid.Grid)
	fmt.Println("Search of word in MxN grid result:", searchInGrid.SearchWord("AAB"))

}
