package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//fmt.Println("arr[2:6] = ", arr[2:6]) // [2 3 4 5]
	//fmt.Println("arr[:6] = ", arr[:6])   // [0 1 2 3 4 5]
	//s1 := arr[2:]
	//fmt.Println("arr[2:] = ", s1) // [2 3 4 5 6 7]
	//s2 := arr[:]
	//fmt.Println("arr[:] = ", s2) // [0 1 2 3 4 5 6 7]
	//
	//fmt.Println("After updateSlice(s1)")
	//updateSlice(s1)
	//fmt.Println(s1)  // [100 3 4 5 6 7]
	//fmt.Println(arr) // [0 1 100 3 4 5 6 7]
	//
	//fmt.Println("After updateSlice(s2)")
	//updateSlice(s2)
	//fmt.Println(s2)  // [100 1 100 3 4 5 6 7]
	//fmt.Println(arr) // [100 1 100 3 4 5 6 7]
	//
	//fmt.Println("Reslice")
	//fmt.Println(s2) // [100 1 100 3 4 5 6 7]
	//s2 = s2[:5]
	//fmt.Println(s2) // [100 1 100 3 4]
	//s2 = s2[2:]
	//fmt.Println(s2) // [100 3 4]

	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1)) // 2, 3, 4, 5
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d", s2, len(s2), cap(s2))   // 5, 6

}
