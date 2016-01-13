// Queue data structure

package queue

type Queue struct {
	Nodes []Node
	first Node
	last  Node
}

type Node struct {
	Item interface{}
	Next *Node
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	return len(q.Nodes)
}

// Empty returns whether the queue is empty.
func (q *Queue) Empty() bool {
	return q.Size() == 0
}

// Enqueue adds an item to the queue.
func (q *Queue) Enqueue(item interface{}) {
	oldlast := q.last
	q.last = Node{
		Item: item,
		Next: nil,
	}
	if q.Empty() {
		q.first = q.last
	} else {
		oldlast.Next = &q.last
	}
	q.Nodes = append(q.Nodes, q.last)
}

// Dequeue removes the least recently added item from the queue.
func (q *Queue) Dequeue() (item interface{}) {
	f := q.first
	item = f.Item
	f = *f.Next
	if q.Empty() {
		//TODO: test whether q.Size() stills returns correctly with a nil Node.
		q.last = Node{nil, nil}
	}
	return item
}
