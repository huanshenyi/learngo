package main

import (
	"main/crawler/engine"
	"main/crawler/paiza/parser"
	"main/crawler/scheduler"
)

func main() {

	//engine.SimpleEngine.Run(engine.Request{
	//	Url:        "https://paiza.jp/career/job_offers/dev_language/Kotlin",
	//	ParserFunc: parser.ParseLanguageList,
	//})
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "https://paiza.jp/career/job_offers/dev_language/Kotlin",
		ParserFunc: parser.ParseLanguageList,
	})
}
