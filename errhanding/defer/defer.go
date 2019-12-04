package main

import (
	"bufio"
	"fmt"
	"log"
	"main/functional/fib"
	"os"
	"strconv"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("panic too many")
		}
	}
}

func writeFile(fileName string) {
	//file, err := os.Create(fileName)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//wrier := bufio.NewWriter(file)
	//f := fib.Fibonacci()
	//for i := 0; i < 20; i++ {
	//	fmt.Fprintln(wrier, f())
	//}
	//wrier.Flush()

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	xx := bufio.NewWriter(file)
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		xx.WriteString(fmt.Sprintln(strconv.Itoa(f())))
		//fmt.Fprintln(xx, f())
	}
	xx.Flush()
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
