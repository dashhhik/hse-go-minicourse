package main

func BinarySearch(n []int, x int) int {
	low, high := 0, len(n)-1
	for low <= high {
		mid := (low + high) / 2
		if n[mid] == x {
			return mid
		} else if n[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
