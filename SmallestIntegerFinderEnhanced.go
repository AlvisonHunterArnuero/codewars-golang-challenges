package kata

import (
	"sort"
)

func SmallestIntegerFinderEnhanced(numbers []int) int {
	sort.Ints(numbers[:])
	return numbers[0]
}
