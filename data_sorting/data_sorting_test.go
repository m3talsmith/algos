package dataSorting

import (
	"reflect"
	"testing"
)

func TestFindDuplicatesNaive(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 1}
	knownDuplicates := []int{1, 5}
	duplicates := FindDuplicatesNaive(numbers)
	if len(duplicates) == 0 {
		t.Fatalf("No duplicates were found")
	} else if !reflect.DeepEqual(duplicates, knownDuplicates) {
		t.Fatalf("\nExpected: %v\nGot: %v", knownDuplicates, duplicates)
	}
}

func TestFindDuplicatesNaiveNoDuplicates(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	duplicates := FindDuplicatesNaive(numbers)
	if len(duplicates) > 0 {
		t.Fatalf("No duplicates should have been found: %v", duplicates)
	}
}
