package middle

func FindMinMax(nums []int) (int, int) {
	minimum := nums[0]
	maximum := nums[0]
	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
		if num > maximum {
			maximum = num
		}
	}
	return minimum, maximum
}
