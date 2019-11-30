package main

import (
	"fmt"
	"main/retriever/mock"
	real2 "main/retriever/real"
	"time"
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
	inspect(r)
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	realRetriever, ok := r.(*real2.Retriever)
	if ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a mock retriver")
	}

	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type switch: ")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Content:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
