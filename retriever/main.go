package main

import (
	"fmt"
	"main/retriever/mock"
	real2 "main/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"mock imooc"}
	r = real2.Retriever{}
	fmt.Println(download(r))
}
