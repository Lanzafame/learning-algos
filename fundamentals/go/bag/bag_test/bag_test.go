package bag_test

import (
	"fmt"
	"testing"

	"github.com/Lanzafame/learning-algos/fundamentals/go/bag"
)

var bagIsEmptyTests = []struct {
	bag      bag.Bag // input
	expected bool    // expected result
}{
	{bag.Bag{}, true},
	{bag.Bag{[]bag.Node{{1, nil}}}, false},
}

var bagSizeTests = []struct {
	bag      bag.Bag // input
	expected int     // expected result
}{
	{bag.Bag{}, 0},
	{bag.Bag{[]bag.Node{{1, nil}}}, 1},
}

var bagAddTests = []struct {
	bag      bag.Bag // input
	node     int     // node to add to bag
	expected bag.Bag // expected result
}{
	{bag.Bag{}, 1, bag.Bag{[]bag.Node{{1, nil}}}},
	{bag.Bag{}, 2, bag.Bag{[]bag.Node{{2, nil}}}},
}

func TestIsEmpty(t *testing.T) {
	for _, tt := range bagIsEmptyTests {
		actual := tt.bag.IsEmpty()
		if actual != tt.expected {
			t.Errorf("Bag.IsEmpty: expected %t, actual %t", tt.expected, actual)
		}
	}
}

func TestSize(t *testing.T) {
	for _, tt := range bagSizeTests {
		actual := tt.bag.Size()
		if actual != tt.expected {
			t.Errorf("Bag.Size: expected %d, actual %t", tt.expected, actual)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, tt := range bagAddTests {
		tt.bag.Add(tt.node)
		if nodeSliceEq(tt.bag.Nodes, tt.expected.Nodes) {
			t.Errorf("Bag.Add(%+v): expected %+v, actual %+v", tt.node, tt.expected, tt.bag)
		}
	}
}

func nodeSliceEq(a, b []bag.Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		fmt.Printf("a[%d]: %+v, b[%d]: %+v", i, a[i], i, b[i])
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
