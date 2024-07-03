package middle

func LinearSearch(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1

}
