package main

func Intersection(n, m []int) []int {
	set := make(map[int]struct{})
	for _, v := range n {
		set[v] = struct{}{}
	}

	var intersection []int
	for _, v := range m {
		if _, found := set[v]; found {
			intersection = append(intersection, v)
			delete(set, v)
		}
	}

	return intersection
}
