package main

import (
	"fmt"
	"main/infra"
)

func getRetriever() retrieve {
	return infra.Retriever{}
}

// ?: Something that can "Get"

type retrieve interface {
	Get(string) string
}

func main() {
	var r retrieve = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
