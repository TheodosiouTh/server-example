package calculator

func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + Sum(numbers[1:])
}
