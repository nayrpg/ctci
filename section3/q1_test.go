package section3

import (
	"testing"

	"github.com/nayrpg/ctci/test"
)

func TestMultiStackTriple(t *testing.T) {
	multi := NewMultiStack[int](3)
	test.Assert(t, multi.IsEmpty(0) && multi.IsEmpty(1) && multi.IsEmpty(2))
	multi.Push(1, 1)
	test.Assert(t, multi.IsEmpty(0) && !multi.IsEmpty(1) && multi.IsEmpty(2))
	test.Assert(t, multi.Peek(1) == 1)
	multi.Push(0, 3)
	test.Assert(t, !multi.IsEmpty(0) && !multi.IsEmpty(1) && multi.IsEmpty(2))
	test.Assert(t, multi.Peek(0) == 3)
	test.Assert(t, multi.Pop(0) == 3)
	test.Assert(t, multi.IsEmpty(0) && !multi.IsEmpty(1) && multi.IsEmpty(2))
	multi.Push(2, 4)
	test.Assert(t, multi.IsEmpty(0) && !multi.IsEmpty(1) && !multi.IsEmpty(2))
}
