package main

import (
	"algos/data_sorting"
	"fmt"
)

func main() {
	if duplicates := dataSorting.FindDuplicatesNaive([]int{3, 1, 11, 2, 4, 3}); duplicates != nil {
		fmt.Println(duplicates)
	}
}
