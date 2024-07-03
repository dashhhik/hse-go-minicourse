package main

func Merge(n, m []int) []int {
	merged := make([]int, 0, len(n)+len(m))
	i, j := 0, 0

	for i < len(n) && j < len(m) {
		if n[i] < m[j] {
			merged = append(merged, n[i])
			i++
		} else {
			merged = append(merged, m[j])
			j++
		}
	}

	for i < len(n) {
		merged = append(merged, n[i])
		i++
	}

	for j < len(m) {
		merged = append(merged, m[j])
		j++
	}

	return merged
}
