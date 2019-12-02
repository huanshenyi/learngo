package main

import (
	"fmt"
	"io"
	"main/retriever/mock"
	real2 "main/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func Post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	io.ReadCloser()
	var r Retriever
	retriever := mock.Retriever{"this is a fack imooc.com"}
	r = retriever
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
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
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
