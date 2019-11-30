package main

import (
	"fmt"
	"main/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())     // 1
	fmt.Println(q.Pop())     // 2
	fmt.Println(q.IsEmpty()) // false
	fmt.Println(q.Pop())     // 3
}