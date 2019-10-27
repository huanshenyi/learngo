package main

import "fmt"

var (
	aa = 3
	ss = "kkk"
	bb = true
	)

func variable()  {
	var a int
	var s string
	fmt.Printf("%d %q",a,s)
}

func variableInitiaValue()  {
	var a, b int = 3,4
	var s string = "abc"
	fmt.Println(a,s,b)
}

func variableTypeDefuction()  {
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func variableShorter()  {
	a,b,c,s := 3,4,true,"def"
	b = 5
	fmt.Println(a,b,c,s)
}

func main()  {
	fmt.Println("hello world")
	variable()
	variableInitiaValue()
	variableTypeDefuction()
	variableShorter()
	fmt.Println(aa,ss,bb)
}
