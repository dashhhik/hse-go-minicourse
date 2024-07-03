package main

func Fibb(n int) []int {
	fibb := []int{0, 1}
	for i := 2; i < n; i++ {
		fibb = append(fibb, fibb[i-1]+fibb[i-2])
	}
	return fibb
}
