package main 



// Queue our own specify implementation
type Queue struct {
	queue []*Node
	lastPos int
	current int
}

// NewQueue create a new Queue object
func NewQueue(size, lastPos int, current int) *Queue {
	queue := make([]*Node, size)
	return &Queue{queue: queue, lastPos: lastPos, current: current}
}

func (q *Queue) add(node *Node){
	q.queue[q.lastPos] = node
	q.lastPos++
}

func (q *Queue) pop() *Node{
	q.current++
	return q.queue[q.current-1]
}