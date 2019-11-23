package main

import "fmt"

func main() {
	var a [3]int
	for i := range a {
		fmt.Printf("a[%d]:%d\n", i, a[i])
	}
}
