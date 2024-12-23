package initialize

import (
	"app/cmd/video-hls-service/queue"
	"sync"
)

func Run() {
	list := []func(){
		runHttpSrver,
		queue.InitQueue,
	}

	var wg sync.WaitGroup

	for _, f := range list {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(f)
	}

	wg.Wait()
}
