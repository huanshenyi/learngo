package parser

import (
	"fmt"
	"main/crawler/engine"
	"regexp"
)

const languageListRe = `<li>
<a href="(/career/job_offers/dev_language/[0-9a-zA-Z]+)">([^<]+)</a>
</li>`

func ParseLanguageList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(languageListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := fmt.Sprintf("%v%s", "https://paiza.jp/", m[1])
		// 要素をitemに追加
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(url),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
