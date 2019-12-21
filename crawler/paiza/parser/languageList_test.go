package parser

import (
	"main/crawler/fetcher"
	"testing"
)

func TestParseLanguageList(t *testing.T) {

	//contents, err := ioutil.ReadFile("languageList_test_data.html")
	contents, err := fetcher.Fetch("https://paiza.jp/career/job_offers/dev_language/Kotlin")
	//fmt.Printf("%s", contents)
	if err != nil {
		panic(err)
	}
	result := ParseLanguageList(contents)
	const resultSize = 11
	expectedUrls := []string{
		"https://paiza.jp//career/job_offers/dev_language/Java", "https://paiza.jp//career/job_offers/dev_language/C", "https://paiza.jp//career/job_offers/dev_language/PHP",
	}
	expectedLanguages := []string{
		"language:Java", "language:C", "language:PHP",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should hava %d requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should hava %d requests; but had %d", resultSize, len(result.Items))
	}

	for i, languages := range expectedLanguages {
		if result.Items[i].(string) != languages {
			t.Errorf("expected languages #%d: %s; but was %s", i, languages, result.Items[i].(string))
		}
	}
}
