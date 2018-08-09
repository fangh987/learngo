package main

import (
	"learngo/crawler/engine"
		"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler/persist"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount:100,
		ItemChan:persist.ItemSaver(),
	}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
e.Run(engine.Request{
	Url:"http://www.zhenai.com/zhenghun/shanghai",
	ParserFunc:parser.ParseCity,
})
}
