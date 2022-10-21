package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	posTotal := 0
	negTotal := 0
	if len(numbers) == 0 {
		return []int{}
	}
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			posTotal += 1
		} else {
			negTotal += numbers[i]
		}
	}
	res = append(res, posTotal, negTotal)
	return res
}
