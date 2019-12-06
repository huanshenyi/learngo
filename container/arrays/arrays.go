package arrays

import "fmt"

func printArray(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
	arr[0] = 100
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println("printArray(arr1)")
	printArray(arr1[:])
	fmt.Println("printArray(arr1)")
	printArray(arr3[:])
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
