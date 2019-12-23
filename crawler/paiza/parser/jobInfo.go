package parser

import (
	"main/crawler/engine"
	"main/crawler/model"
	"regexp"
)

var incomeRe = regexp.MustCompile(`<div class='strong font18 color_blue'>([^<]+)</div>`)
var companyNameRe = regexp.MustCompile(`<h2 class='ttl mt0 mb0' id='corpname'>([^<]+)</h2>`)
var jobInfoRe = regexp.MustCompile(`<h1 class='ttl mt0 mb0'>([^<]+)</h1>`)
var positionRe = regexp.MustCompile(`<td class='font16' colspan='3'>
<strong>([^<]+)</strong>
</td>`)

func ParserJobInfo(contents []byte) engine.ParseResult {
	jobInfo := model.JonInfo{}
	// 提示年収
	jobInfo.Income = extractString(contents, incomeRe)
	// 会社名
	jobInfo.CompanyName = extractString(contents, companyNameRe)
	// 仕事内容
	jobInfo.JobInfo = extractString(contents, jobInfoRe)
	// ポジション
	jobInfo.Position = extractString(contents, positionRe)

	result := engine.ParseResult{
		Items: []interface{}{jobInfo},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
