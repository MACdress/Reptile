package persist

import "github.com/astaxie/beego"

func ItemSaver() chan interface{}{
	out:= make(chan interface{})
	go func(){
		for {
			item := <-out
			beego.Info(item)
		}
	}()
	return out
}
