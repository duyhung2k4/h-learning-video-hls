package queue

import (
	"app/config"
	"app/constant"
	"app/dto/queuepayload"
	"app/service"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type queueFileM3U8 struct {
	videoService service.VideoService
}

type QueueFileM3U8 interface {
	Worker()
}

func (q *queueFileM3U8) Worker() {
	conn := config.GetRabbitmq()
	ch, err := conn.Channel()
	if err != nil {
		log.Println("error chanel: ", err)
		return
	}

	qu, err := ch.QueueDeclare(
		string(constant.QUEUE_FILE_M3U8),
		true,
		false,
		false,
		false,
		amqp091.Table{},
	)
	if err != nil {
		log.Println("error queue declare: ", err)
		return
	}
	log.Printf("start %s", string(constant.QUEUE_FILE_M3U8))

	msgs, err := ch.Consume(
		qu.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("error consumer: ", err)
		return
	}

	for d := range msgs {
		go func(mess amqp091.Delivery) {
			var payload queuepayload.QueueFileM3U8Payload
			err := json.Unmarshal(mess.Body, &payload)
			if err != nil {
				log.Println("error msg: ", err)
				mess.Reject(true)
				return
			}

			listFile, err := q.videoService.GetListVideo(payload)
			if err != nil {
				log.Println("error get list file: ", err)
				mess.Reject(true)
				return
			}

			err = q.videoService.DownloadVideo(listFile, payload)
			if err != nil {
				log.Println("error download video: ", err)
				mess.Reject(true)
				return
			}

			mess.Ack(false)
		}(d)
	}
}

func NewQueueFileM3U8() QueueFileM3U8 {
	return &queueFileM3U8{
		videoService: service.NewVideoService(),
	}
}
