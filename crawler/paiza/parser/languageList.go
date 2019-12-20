package parser

import (
	"fmt"
	"main/crawler/engine"
	"regexp"
)

const languageListRe = `<a href="(/career/job_offers/dev_language/[0-9a-zA-Z]+)">([^<]+)</a>`

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
	//同じものを省き
	//result = RemoveRepByLoop(result)
	return result
}

//func RemoveRepByLoop(slc engine.ParseResult) engine.ParseResult {
//	var result engine.ParseResult  // 存放结果
//	for i := range slc.Requests{
//		flag := true
//		for j := range result.Requests{
//			if slc.Items[i] == result.Items[j] {
//				flag = false  // 存在重复元素，标识为false
//				break
//			}
//		}
//		if flag {  // 标识为false，不添加进结果
//			result = append(result, slc[i])
//		}
//	}
//	return result
//}
