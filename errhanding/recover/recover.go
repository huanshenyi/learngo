package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is an error"))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	f := fmt.Errorf("%v", "err hihi")
	panic(f) //Error occurred: err hihi
}

func main() {
	tryRecover()
}
