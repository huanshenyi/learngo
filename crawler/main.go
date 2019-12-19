package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding"

	"golang.org/x/net/html/charset"
)

func main() {
	resp, err := http.Get("https://paiza.jp/career/job_offers/dev_language/Java")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s", all)
	printCityList(all)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	// htmlのcharsetを予測
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(/career/job_offers/dev_language/[0-9a-zA-Z]+)">([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		url := fmt.Sprintf("%v%s", "https://paiza.jp/", m[1])
		fmt.Printf("Language:%s, URL: %s\n", m[2], url)
	}
	fmt.Printf("Matches found:%d\n", len(matches))
}
