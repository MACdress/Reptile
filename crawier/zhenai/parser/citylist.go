package parser

import (
	"Reptile/crawier/engine"
	"regexp"
)

var  CityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func CityList(contents []byte) engine.ParseResult{
	result := engine.ParseResult{}
	maces := CityListRe.FindAllSubmatch(contents,-1)
	for _,m := range maces {
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCity,
		})
	}
	return result
}
