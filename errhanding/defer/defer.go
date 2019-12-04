package main

import (
	"bufio"
	"errors"
	"fmt"
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

	//file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error")
	if err != nil {
		//log.Fatal(err.Error())
		//fmt.Println("Error:", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%v,%v,%v", pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
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
