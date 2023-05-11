//* Stack
//* ref : https://www.youtube.com/watch?v=qv_7cZ-hUXc

package main

import "fmt"

type Stack struct {
	Data []int
	Size int
	Top  int
}

func NewStack(size int) *Stack {
	return &Stack{
		Data: make([]int, size),
		Size: size,
		Top:  -1,
	}
}

func (s *Stack) Push(x int) {
	if s.Top == s.Size-1 {
		fmt.Println("error, stackoverflow")
	} else {
		s.Top++
		s.Data[s.Top] = x
		fmt.Println(s.Data, s.Top)
	}
}

func (s *Stack) Pop() int {
	var x int
	if s.Top == -1 {
		fmt.Println("error, stackunderflow")
	} else {
		x = s.Data[s.Top]
		s.Data[s.Top] = 0
		s.Top--
		fmt.Println(s.Data, s.Top)
	}
	return x
}

func (s *Stack) Peek() int {
	var x int
	if s.Data[s.Top] == 0 {
		fmt.Println("error, invalid position")
	} else {
		x = s.Data[s.Top]
		fmt.Println(s.Data, s.Top)
	}
	return x
}

func main() {
	s := NewStack(4)

	fmt.Println(s)

	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Pop()
	s.Pop()
	s.Peek()
}
