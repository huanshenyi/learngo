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
const worklocationRe = `<td class='c-job_offer-detail__description'>([^<]+)</td>`

func ParseLanguage(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(languageRe)
	matches := re.FindAllSubmatch(contents, -1)
	re2 := regexp.MustCompile(worklocationRe)
	matches2 := re2.FindAllStringSubmatch(string(contents), -1)
	result := engine.ParseResult{}
	for k, m := range matches {
		url := fmt.Sprintf("%v%s", "https://paiza.jp/", m[1])
		// 要素をitemに追加
		result.Items = append(result.Items, "Job "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParserJobInfo(c, matches2[k][1])
			},
		})
	}
	return result
}
