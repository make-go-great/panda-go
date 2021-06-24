package stack

import "fmt"

type StackInt struct {
	arr []int
}

func NewStackInt() *StackInt {
	return &StackInt{
		arr: make([]int, 0, defaultStackLen),
	}
}

func (s *StackInt) Push(v int) {
	s.arr = append(s.arr, v)
}

func (s *StackInt) Peek() (int, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}

	result := s.arr[len(s.arr)-1]

	return result, true
}

func (s *StackInt) Pop() (int, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}

	result := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return result, true
}

func (s *StackInt) Print() {
	fmt.Println(s.arr)
}
