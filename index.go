package main

import (
	"app/config"
	"app/queue"
	"app/router"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		server := &http.Server{
			Addr:           ":" + config.GetAppPort(),
			Handler:        router.AppRouter(),
			MaxHeaderBytes: 1 << 20,
		}

		log.Fatalln(server.ListenAndServe())
	}()

	go func() {
		defer wg.Done()
		queue.InitQueue()
	}()

	wg.Wait()
}
