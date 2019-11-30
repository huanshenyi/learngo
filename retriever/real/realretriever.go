package real

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	con, _ := ioutil.ReadAll(resp.Body)
	return string(con)
}
