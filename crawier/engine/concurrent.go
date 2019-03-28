package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadNotifier interface {
	WorkerReady(w chan Request)
}

var visitedUrls = make(map[string]bool)

func (c *ConcurrentEngine) Run(seeds ...Request){
	c.Scheduler.Run()
	out:= make(chan ParseResult)
	for i:=0;i<c.WorkerCount;i++{//创建worker等待执行
		createWorker(c.Scheduler.WorkerChan(),out,c.Scheduler)
	}
	for _,r := range  seeds{
		if !(isDuplicate(r.Url)) {
			c.Scheduler.Submit(r)
		}
	}
	for{
		result := <-out
		for _,item := range result.Items{
			go func(){
				c.ItemChan<-item
			}()
		}
		for _,request := range result.Requests{
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult, read ReadNotifier) {
	go func(){
		for {
			read.WorkerReady(in)
			request :=  <-in
			result,err := Worker(request)
			if err!=nil{
				continue
			}
			out <- result
		}
	}()
}


func isDuplicate(url string)bool{
	if visitedUrls[url]{
		return true
	}
	visitedUrls[url] = true
	return false
}