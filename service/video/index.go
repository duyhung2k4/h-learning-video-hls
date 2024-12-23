package videoservice

import (
	"app/internal/connection"
	queuepayload "app/internal/dto/queue_payload"

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

func Register() VideoService {
	return &videoService{
		connRabbitmq: connection.GetRabbitmq(),
	}
}
