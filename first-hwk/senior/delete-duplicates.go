package main

func DeleteDuplicates(nums []int) []int {
	duplicates := make(map[int]struct{})
	newArr := make([]int, 0, len(nums))

	for _, v := range nums {
		if _, exists := duplicates[v]; !exists {
			newArr = append(newArr, v)
			duplicates[v] = struct{}{}
		}
	}

	return newArr
}
