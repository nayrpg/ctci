package section3

type MultiStack[T any] struct {
	tops []int
	arr  []T
}

func NewMultiStack[T any](numStacks int) *MultiStack[T] {
	tops := make([]int, numStacks)
	for i, _ := range tops {
		tops[i] = -1
	}
	return &MultiStack[T]{tops: tops} //Need to initialize with the generic type
}

func (s *MultiStack[T]) Push(stackNum int, val T) { // need set method struct pointer with generic type
	if s.tops[stackNum] < 0 {
		s.tops[stackNum] = stackNum
	} else {
		s.tops[stackNum] += len(s.tops)
	}
	if len(s.arr) <= s.tops[stackNum] {
		s.arr = append(s.arr, make([]T, len(s.tops))...) // need to use spread operator to allow the append
	}
	s.arr[s.tops[stackNum]] = val
}

func (s *MultiStack[T]) Pop(stackNum int) T {
	ret := s.arr[s.tops[stackNum]]
	s.tops[stackNum] -= len(s.tops)
	return ret
}

func (s *MultiStack[T]) Peek(stackNum int) T {
	return s.arr[s.tops[stackNum]]
}

func (s *MultiStack[T]) IsEmpty(stackNum int) bool {
	return s.tops[stackNum] < 0
}
