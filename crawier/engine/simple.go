package engine

import (
	"Reptile/crawier/fetcher"
	"github.com/astaxie/beego"
)

type SimpleEngine struct {

}
func (e *SimpleEngine)Run(seeds ...Request){
	var reqs []Request
	for _,temp := range seeds{
		reqs = append(reqs,temp)
	}
	for len(reqs)>0{
		r:= reqs[0]
		reqs = reqs[1:]
		parseResult,err :=Worker(r)
		if err!=nil{
			beego.Info(err)
			continue
		}
		reqs = append(reqs,parseResult.Requests...)//加三个点的意思在展开这个slice，把里面的东西都加进去
	}
}

func Worker (req Request)(ParseResult,error){
	body,err := fetcher.Fetch(req.Url)
	if err!=nil{
		beego.Info("Fetcher:error,fetcher url:",req.Url)
		return ParseResult{},err
	}
	return req.ParserFunc(body),nil
}
