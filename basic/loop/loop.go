package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.FormatInt(int64(lsb), 10) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	read := bufio.NewReader(file)
	b := make([]byte, 1024)
	n, err := read.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n) //26
	fmt.Println(string(b[:n]))
	/*
		abcde
		12345
		hello
		again
	*/

}

func main() {
	//fmt.Println(
	//	convertToBin(5), //101
	//	convertToBin(13), // 1101
	//	)
	//fmt.Println(strconv.FormatUint(5,2))
	//printFile("abc.txt")
	readFile("abc.txt")
}
