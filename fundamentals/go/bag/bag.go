// Bag

// Methods:
//   - add(interface{})
//   - isEmpty()
//   - size()

package bag

type Bag struct {
	Nodes []Node
}

type Node struct {
	Item interface{}
	Next *Node
}

func (b *Bag) IsEmpty() bool {
	return len(b.Nodes) == 0
}

func (b *Bag) Size() int {
	return len(b.Nodes)
}

func (b *Bag) Add(item interface{}) {
	oldfirst := b.Nodes[0]
	first := Node{}
	first.Item = item
	first.Next = &oldfirst
}
