package queue

func InitQueue() {
	queueFileM3U8 := NewQueueFileM3U8()

	go func() {
		queueFileM3U8.Worker()
	}()
}
