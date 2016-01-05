// Bag

// Methods:
//   - add(interface{})
//   - isEmpty()
//   - size()

package bag

type Bag struct {
	nodes []Node
}

type Node struct {
	item interface{}
	next *Node
}

func (b *Bag) IsEmpty() bool {
	return len(b.nodes) == 0
}

func (b *Bag) Size() int {
	return len(b.nodes)
}

func (b *Bag) Add(item interface{}) {
	oldfirst := b.nodes[0]
	first := Node{}
	first.item = item
	first.next = &oldfirst
}
