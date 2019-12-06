package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q", a, s)
}

func variableInitiaValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDefuction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Println(
		cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

const filename string = "abc.txt"

func consts() {
	const a, b = 3, 4
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(math.Pow(2, 10))
}

func main() {
	fmt.Println("hello world")
	variable()
	variableInitiaValue()
	variableTypeDefuction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
