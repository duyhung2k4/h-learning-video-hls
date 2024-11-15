package service

import (
	"app/config"
	"app/constant"
	"app/dto/queuepayload"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

type videoService struct {
	connRabbitmq *amqp091.Connection
}

type VideoService interface {
	processDownload(filename string, payload queuepayload.QueueFileM3U8Payload) error
	GetListVideo(payload queuepayload.QueueFileM3U8Payload) ([]string, error)
	DownloadVideo(listfile []string, payload queuepayload.QueueFileM3U8Payload) error
}

func (s *videoService) processDownload(filename string, payload queuepayload.QueueFileM3U8Payload) error {
	url := fmt.Sprintf("%s/%s/%s", payload.IpServer, payload.Path, filename)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	dir := fmt.Sprintf("video/%s/%s", payload.Uuid, payload.Quantity)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	filesave := fmt.Sprintf("%s/%s", dir, filename)
	out, err := os.Create(filesave)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *videoService) GetListVideo(payload queuepayload.QueueFileM3U8Payload) ([]string, error) {
	var list []string

	url := fmt.Sprintf("%s/%s", payload.IpServer, payload.Path)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching data: ", err)
		return list, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
		return list, err
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		log.Fatal("Error decoding JSON: ", err)
		return list, err
	}

	return list, nil
}

func (s *videoService) DownloadVideo(listfile []string, payload queuepayload.QueueFileM3U8Payload) error {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var listError []error

	dir := fmt.Sprintf("video/%s/%s", payload.Uuid, payload.Quantity)
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
		"http://%s:%s/api/v1/video/%s/%s/%s_%s0.m3u8",
		config.GetAppHost(),
		config.GetAppPort(),
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

func NewVideoService() VideoService {
	return &videoService{
		connRabbitmq: config.GetRabbitmq(),
	}
}
