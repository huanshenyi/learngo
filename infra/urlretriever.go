package infra

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Retriever struct{}

func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(body)
}
