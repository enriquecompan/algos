package algos

// Bubble Sort implementation
func BubbleSort(array []int) {
	for i := len(array) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

// Selection Sort implementation
func SelectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		min := i
		found := false

		for j := i + 1; j < len(array); j++ {

			if array[j] < array[min] {
				min = j
				found = true
			}
		}

		if found {
			array[i], array[min] = array[min], array[i]
		}
	}
}

// QuickSort implementation
func QuickSort(array []int) []int {

	// Base case
	if len(array) < 2 {
		return array
	}

	pivot := array[0]
	left := []int{}
	right := []int{}
	for i := 1; i < len(array); i++ {
		if array[i] <= pivot {
			left = append(left, array[i])
		} else if array[i] > pivot {
			right = append(right, array[i])
		}
	}

	// Recursive case
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
