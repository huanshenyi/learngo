package queue

import (
	"fmt"
	"testing"
)

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}

func TestQueue_Pop(t *testing.T) {
	list := []Queue{
		{1, 2, 3},
		{4, 5, 7},
		{8, 9, 10},
	}
	for _, tt := range list {
		startNum := tt[0]
		popnum := tt.Pop()
		if popnum != startNum {
			t.Errorf("Pop func error, popnum:%d a[0]:%d", popnum, startNum)
		}
	}
}
