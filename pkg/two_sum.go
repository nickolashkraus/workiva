package pkg

// Array Target Sum
//
// Part 1: Given an int array (int[] arr) and a target value (int target),
// create a function that returns true if the array contains any two values
// whose sum is equal to the target value, false otherwise.
//
// Part 2: Create a method that accepts a list of integers and a target integer
// sum, find and return all unique pairs in the list of integers whose sum is
// equal to the target sum.
//
// NOTES
//   - Use a hash table for complement lookups.
func twoSum1(nums []int, target int) bool {
	table := make(map[int]int)
	for i, num := range nums {
		_, ok := table[target-num]
		if ok {
			return true
		} else {
			table[num] = i
		}
	}
	return false
}

func twoSum2(nums []int, target int) [][]int {
	compls := [][]int{}
	table := make(map[int]int)
	for i, num := range nums {
		_, ok := table[target-num]
		// Cannot use element, since its zero value (0) could be a valid value.
		if ok {
			compls = append(compls, []int{target - num, num})
		} else {
			table[num] = i
		}
	}
	return compls
}
