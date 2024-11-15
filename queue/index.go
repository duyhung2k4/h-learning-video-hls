package queue

import (
	"sync"
)

func InitQueue() {
	queueFileM3U8 := NewQueueFileM3U8()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		queueFileM3U8.Worker()
	}()

	wg.Wait()
}
