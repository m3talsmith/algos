package main

import (
	"fmt"

	"github.com/m3talsmith/algos/data_sorting"
)

func main() {
	if duplicates := dataSorting.FindDuplicatesNaive([]int{3, 1, 11, 2, 4, 3}); duplicates != nil {
		fmt.Println(duplicates)
	}
}
