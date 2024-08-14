package section2

import (
	"slices"
	"testing"
)

func Test1Q4(t *testing.T) {
	head := &LLNode[int]{Val: 1}
	head.InsertAtEnd(2)
	head.InsertAtEnd(1)
	head.InsertAtEnd(3)
	head = head.PartitionOnX(2)
	head.Print()
	assert(t, slices.Equal(head.ToSlice(), []int{1, 1, 2, 3}))
}

func Test2Q4(t *testing.T) {
	head := &LLNode[int]{Val: 2}
	head.InsertAtEnd(4)
	head.InsertAtEnd(5)
	head.InsertAtEnd(1)
	head.InsertAtEnd(6)
	head.InsertAtEnd(1)
	head.InsertAtEnd(2)
	head = head.PartitionOnX(2)
	head.Print()
	assert(t, slices.Equal(head.ToSlice(), []int{1, 1, 2, 4, 5, 6, 2}))
}

func assert(t *testing.T, cond bool) {
	if !cond {
		t.Errorf("condition failed")
	}
}
