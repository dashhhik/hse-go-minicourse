package main

type Queue struct {
	Stack1 []int
	Stack2 []int
}

func (q *Queue) Push(x int) {
	q.Stack1 = append(q.Stack1, x)
}

func (q *Queue) Pop() int {
	if len(q.Stack2) == 0 {
		for len(q.Stack1) > 0 {
			n := len(q.Stack1) - 1
			q.Stack2 = append(q.Stack2, q.Stack1[n])
			q.Stack1 = q.Stack1[:n]
		}
	}
	n := len(q.Stack2) - 1
	x := q.Stack2[n]
	q.Stack2 = q.Stack2[:n]
	return x
}
