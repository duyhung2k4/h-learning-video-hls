package main

import (
	"app/cmd/video-hls-service/initialize"
	"app/internal/connection"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		initialize.Run()
	}()

	wg.Wait()

	connection.DeferFunc()
}
