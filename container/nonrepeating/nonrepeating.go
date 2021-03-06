package main

import (
	"fmt"
)

func LengthOfNonRepeatingSubStr(s string) int {
	var lastOccurred = make(map[byte]int)
	//lastOccurred := make([]int, 0xffff)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		//fmt.Println(i, ch)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			//もしmapに存在すれば,一個前へ進
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			//fmt.Printf("maxLength%d\n", maxLength)
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		LengthOfNonRepeatingSubStr("abbcv"))
}
