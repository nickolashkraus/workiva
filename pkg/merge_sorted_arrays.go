package pkg

// Merge Sorted Arrays
//
// Create a function that merges two sorted arrays into a new sorted array.
func mergeSortedArrays(l1 []int, l2 []int) []int {
	sorted := []int{}
	i, j, k := 0, 0, 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			sorted = append(sorted, l1[i])
			i, k = i+1, k+1
		} else {
			sorted = append(sorted, l2[j])
			j, k = j+1, k+1
		}
	}
	for i < len(l1) {
		sorted = append(sorted, l1[i])
		i, k = i+1, k+1
	}
	for j < len(l2) {
		sorted = append(sorted, l2[j])
		j, k = j+1, k+1
	}
	return sorted
}
