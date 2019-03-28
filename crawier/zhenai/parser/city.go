package parser

import (
	"Reptile/crawier/engine"
	"Reptile/utils"
	"regexp"
)

var  CityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var  CityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
func ParseCity(contents []byte) engine.ParseResult{
	result := engine.ParseResult{}
	maces := CityRe.FindAllSubmatch(contents,-1)
	for _,m := range maces {
		img := getImg(contents,string(m[1]))
		name := string(m[2])
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:func (c []byte)engine.ParseResult{
				return ParseProfile(c ,img,name,string(m[1]))
			},
		})
	}
	matches :=CityUrlRe.FindAllSubmatch(contents,-1)
	for _,m := range matches{
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCity,
		})
	}
	return result
}

func getImg(contents []byte,url string)string {
	var MyImgRe = regexp.MustCompile(`<div class="photo"><a href="`+url+`"[^>]*><img src="(https://photo.zastatic.com/images/photo/[^\?]+)\?[^>]*></a>`)
	myImg := utils.ExtractString(contents, MyImgRe)
	return myImg
}

