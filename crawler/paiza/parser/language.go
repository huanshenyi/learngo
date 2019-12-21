package parser

import (
	"fmt"
	"main/crawler/engine"
	"regexp"
)

const languageRe = `<a class="c-job_offer-box__header__title__link" href="(/career/job_offers/[1-9]+\?from=list)"><h3 class='c-job_offer-box__header__title'>
([^<]+)
</h3>
</a>`

func ParseLanguage(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(languageRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := fmt.Sprintf("%v%s", "https://paiza.jp/", m[1])
		// 要素をitemに追加
		result.Items = append(result.Items, "Job "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(url),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
