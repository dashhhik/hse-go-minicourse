package main

func CountElement(n []int, x int) int {
	count := 0
	for i := 0; i < len(n); i++ {
		if n[i] == x {
			count++
		}
	}
	return count
}
