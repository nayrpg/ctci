package section2

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type LLNode[T constraints.Ordered] struct {
	Val  T
	Next *LLNode[T]
}

func (head *LLNode[T]) Print() {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%v", cur.Val)
		if cur.Next != nil {
			fmt.Printf(" -> ")
		}
	}
}

func (head *LLNode[T]) InsertAtEnd(val T) {
	if head == nil {
		panic("called on nil head")
	}
	cur := head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = &LLNode[T]{Val: val}
}

func (head *LLNode[T]) ToSlice() []T {
	slice := make([]T, 0)
	for cur := head; cur != nil; cur = cur.Next {
		slice = append(slice, cur.Val)
	}
	return slice
}

func (head *LLNode[T]) PartitionOnX(x T) *LLNode[T] {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead, newTail *LLNode[T]
	// : needed for var assignment and for loop should run when cur.Next isn't
	// nil
	for cur := head; cur.Next != nil; {
		if cur.Next.Val < x { // should be the Val of cur.Next
			if newHead == nil {
				newHead = cur.Next
				newTail = newHead
			} else {
				newTail.Next = cur.Next
				newTail = newTail.Next
			}
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	if newTail != nil {
		newTail.Next = head
		return newHead
	}
	return head
}
