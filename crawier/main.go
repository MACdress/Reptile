package main

import (
	"Reptile/crawier/engine"
	"Reptile/crawier/zhenai/parser"
	"Reptile/persist"
	"Reptile/scheduler"
)

func main(){
	e:= engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan : persist.ItemSaver(),
	}
	e.Run(
		engine.Request{
			Url:`http://www.zhenai.com/zhenghun`,
			ParserFunc:parser.CityList,
		})
	//e.Run(
	//	engine.Request{
	//		Url:"http://www.zhenai.com/zhenghun/shanghai",
	//		ParserFunc:parser.ParseCity,
	//	})
}