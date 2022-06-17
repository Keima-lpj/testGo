package queue

type Queue []int

type Queue2 [][]int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue2) Push(v []int) {
	*q = append(*q, v)
}

func (q *Queue2) Pop() []int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue2) IsEmpty() bool {
	return len(*q) == 0
}
