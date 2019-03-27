package parser

import (
	"Reptile/crawier/engine"
	"Reptile/models/user"
	"Reptile/utils"
	"github.com/astaxie/beego"
	"regexp"
	"strconv"
)


var AgeRe  = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)//年龄
var NameRe = regexp.MustCompile(`<div class="info" data-v-5b109fc3><div class="name" data-v-5b109fc3><h1 class="nickName" data-v-5b109fc3>([^<]+)</h1>`)//姓名
var HeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)//身高
var WeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)//体重
var IncomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)//月收入
var MarriageRe = regexp.MustCompile(`<div class="purple-btns" data-v-bff6f798><div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)//婚姻状况
var OccupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>工作地:([^<]+)</div>`)//工作地
var XinZuoRe = regexp.MustCompile(`岁</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)//星座
var OtherImgRe = regexp.MustCompile(`(https://photo.zastatic.com/images/photo/[^\?]+)\?`)//照片

func ParseProfile(contents []byte,img string ,name string)engine.ParseResult{
	profile := user.Profile{
		MyImg:img,
		Name:name,
	}
	if age,err := strconv.Atoi(utils.ExtractString(contents,AgeRe));err!=nil {
		beego.Info("年龄值不正确")
	}else {
		profile.Age = age
	}
	marriage := utils.ExtractString(contents,MarriageRe)
	profile.Marriage = marriage
	profile.Name = utils.ExtractString(contents,NameRe)
	profile.OtherImg = extractImg(contents)
	if height,err := strconv.Atoi(utils.ExtractString(contents,HeightRe));err!=nil{
		beego.Info("身高值不正确")
	}else{
		profile.Height = height
	}
	if weight,err := strconv.Atoi(utils.ExtractString(contents,WeightRe));err!=nil{
		beego.Info("体重值不正确")
	}else{
		profile.Weight = weight
	}
	profile.Income = utils.ExtractString(contents,IncomeRe)
	profile.Occupation = utils.ExtractString(contents,OccupationRe)
	profile.XinZuo = utils.ExtractString(contents,XinZuoRe)
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}


func extractImg (contents []byte)[]string{
	var imgs []string
	maces := OtherImgRe.FindAllSubmatch(contents,-1)
	for _,m := range maces {
		imgs = append(imgs,string(m[1]))
	}
	return imgs
}