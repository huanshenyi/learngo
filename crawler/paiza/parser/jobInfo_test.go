package parser

import (
	"io/ioutil"
	"main/crawler/model"
	"testing"
)

func TestParserJobInfo(t *testing.T) {
	contents, err := ioutil.ReadFile("jobInfo_test_data.html")
	if err != nil {

	}
	result := ParserJobInfo(contents, "東京都豊島区池袋 2-9-4")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	jonInfo := result.Items[0]
	expected := model.JonInfo{
		CompanyName:  "株式会社EPARKグルメ",
		Income:       "350万 〜 800万円",
		JobInfo:      "Androidアプリエンジニア◎外食事業に特化した便利サービスを開発【フレックス/月残業平均15時間以下】",
		Language:     "Kotlin",
		Worklocation: "東京都豊島区池袋 2-9-4",
	}
	if jonInfo != expected {
		t.Errorf("expected %v; but was %v", expected, jonInfo)
	}
}
