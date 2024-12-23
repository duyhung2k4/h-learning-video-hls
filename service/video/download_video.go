package videoservice

import (
	"app/internal/connection"
	constant "app/internal/constants"
	queuepayload "app/internal/dto/queue_payload"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

func (s *videoService) DownloadVideo(listfile []string, payload queuepayload.QueueFileM3U8Payload) error {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var listError []error

	dir := fmt.Sprintf("cmd/video-hls-service/data/video/%s/%s", payload.Uuid, payload.Quantity)
	os.RemoveAll(dir)

	for _, f := range listfile {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			err := s.processDownload(filename, payload)
			if err != nil {
				mtx.Lock()
				listError = append(listError, err)
				mtx.Unlock()
			}
		}(f)
	}

	wg.Wait()

	if len(listError) > 0 {
		for i, e := range listError {
			log.Printf("error download file video %d: %s", i, e.Error())
		}
		return errors.New("error download file video")
	}

	ch, err := s.connRabbitmq.Channel()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"http://%s:%s/api/v1/video/%s/%s/%s_%s.m3u8",
		connection.GetConnect().VideoHlsService.Host,
		connection.GetConnect().VideoHlsService.Port,
		payload.Uuid,
		payload.Quantity,
		payload.Uuid,
		payload.Quantity,
	)
	payloadMess := queuepayload.QueueUrlQuantityPayload{
		Url:      url,
		Quantity: payload.Quantity,
		Uuid:     payload.Uuid,
	}
	payloadJsonString, err := json.Marshal(payloadMess)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(context.Background(),
		"",
		string(constant.QUEUE_URL_QUANTITY),
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        payloadJsonString,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
