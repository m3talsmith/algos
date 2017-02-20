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
 * checking the for a duplicate. If there's a duplicate, we advance the right
 * neighbor index until it no longer matches. Once we have gone through the
 * whole list we'll return the list of possible duplicates.
 */
func FindDuplicatesNaive(numbers []int) []int {
	i := 0
	count := len(numbers)
	duplicates := []int{}
	sortedNumbers := sort.IntSlice(numbers)
	if !sort.IntsAreSorted(sortedNumbers) {
		sortedNumbers.Sort()
	}

	var ni int
	for i < count {
		ni = i + 1
		if ni >= count {
			i = ni
			break
		}
		if sortedNumbers[i] == sortedNumbers[ni] {
			duplicates = append(duplicates, sortedNumbers[i])
			// Increment the right neighbor index until it no longer matches
			for sortedNumbers[i] == sortedNumbers[ni] {
				ni = ni + 1
			}
		}
		i = ni
	}

	return duplicates
}
