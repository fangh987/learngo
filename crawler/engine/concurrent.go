package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	//初始化
	e.Scheduler.Run()
	//创建多个线程通道
	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//反复循环查找所有
	for {
		result := <-out
		for _, item := range result.Item {
			go func() {e.ItemChan <- item}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

//请求解析返回out储存
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i`m ready
			ready.WorkerReady(in)
			//当in为空时，该线程会停止在此处，等到chan Request有值时，会继续向下执行（无关乎方法和参数，只要此chan Request有值即可）
			request := <-in
			//请求路径解析，返回对应页面数据
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
