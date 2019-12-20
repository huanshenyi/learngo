package main

import (
	"main/crawler/engine"
	"main/crawler/paiza/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "https://paiza.jp/career/job_offers/dev_language/Kotlin",
		ParserFunc: parser.ParseLanguageList,
	})

}
