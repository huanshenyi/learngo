package parser

import (
	"main/crawler/engine"
	"main/crawler/model"
	"regexp"
)

var incomeRe = regexp.MustCompile(`<div class='strong font18 color_blue'>([^<]+)</div>`)
var companyNameRe = regexp.MustCompile(`<h2 class='ttl mt0 mb0' id='corpname'>([^<]+)</h2>`)
var jobInfoRe = regexp.MustCompile(`<h1 class='ttl mt0 mb0'>([^<]+)</h1>`)
var LanguageRe = regexp.MustCompile(`<td><ul class="clearfix mb0"><li class="lang_tag font14 priority"><a href="/career/job_offers/dev_language/[a-zA-Z]+">([^<]+)</a></li>`)

func ParserJobInfo(contents []byte, worklocation string) engine.ParseResult {
	jobInfo := model.JonInfo{}
	// 勤務地
	jobInfo.Worklocation = worklocation
	// 提示年収
	jobInfo.Income = extractString(contents, incomeRe)
	// 会社名
	jobInfo.CompanyName = extractString(contents, companyNameRe)
	// 仕事内容
	jobInfo.JobInfo = extractString(contents, jobInfoRe)
	// ポジション
	jobInfo.Language = extractString(contents, LanguageRe)

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
