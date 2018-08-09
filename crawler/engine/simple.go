package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, items := range parseResult.Item {
			log.Printf("Got item %v", items)
		}

	}
}

//获取路径对应的页面信息
func Worker(r Request) (ParseResult, error) {
	log.Printf("Fentch %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher ：error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
