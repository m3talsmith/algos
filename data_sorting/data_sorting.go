package dataSorting

import "sort"

// FindDuplicatesNaive returns any duplicates in slice of integers
/*
 * Hypothesis
 * -------
 * If we sort the integers then check the neighbor to the right, we'll know if
 * there's a duplicate. If there is a duplicate, we can append that to the
 * list of previously found duplicates, increase the index by two, and
 * continue. If there is no duplicate, we increase the index by one and
 * continue. We will have to check that the right isn't out of bounds before
 * checking the for a duplicate. Once we have gone through the whole list
 * we'll return the list of possible duplicates.
 *
 * Possible Problems
 * -----------------
 * It is possible that we find a right neighbor that is a duplicate, but fail
 * to account for any duplicates to the right of that neighbor. It's not to
 * then loop over that, but the duplicate has already been marked. From a
 * naive implementation, this work though.
 */
func FindDuplicatesNaive(numbers []int) []int {
	i := 0
	count := len(numbers)
	duplicates := []int{}
	sortedNumbers := sort.IntSlice(numbers)
	if !sort.IntsAreSorted(sortedNumbers) {
		sortedNumbers.Sort()
	}

	var currentInt, rightInt, nextInt int
	for i < count {
		nextInt = i + 1
		if nextInt >= count {
			i = nextInt
			break
		}
		currentInt = sortedNumbers[i]
		rightInt = sortedNumbers[nextInt]
		if currentInt == rightInt {
			duplicates = append(duplicates, currentInt)
			i = i + 2
		} else {
			i = nextInt
		}
	}
	return duplicates
}
