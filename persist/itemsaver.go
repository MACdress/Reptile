package persist

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
)

func ItemSaver() chan interface{}{
	out:= make(chan interface{})
	go func(){
		for {
			item := <-out
			beego.Info(item)
			id,err:= save(item)
			if err!= nil{
				beego.Info("插入失败")
			}
			beego.Info("Id = ",id)
		}
	}()
	return out
}

func save(item interface{})(id string,err error) {
	 client,err := elastic.NewClient(elastic.SetSniff(false))
	 if err!=nil {
	 	beego.Info("创建ES客户端失败")
	 	return "",err
	 }
	 resp,err := client.Index().
	 	Index(beego.AppConfig.String("Es_name")).
	 	Type(beego.AppConfig.String("Es_table")).
	 	BodyJson(item).
	 	Do(context.Background())
	 if err !=nil{
	 	return "",err
	 }
	 return resp.Id,nil
}
