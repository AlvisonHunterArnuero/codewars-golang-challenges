package kata

func SmallestIntegerFinder(numbers []int) int {
	min := numbers[0]
	for _, value := range numbers {
		if value < min {
			min = value
		}
	}
	return min
}
