package videoservice

import (
	queuepayload "app/internal/dto/queue_payload"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (s *videoService) processDownload(filename string, payload queuepayload.QueueFileM3U8Payload) error {
	url := fmt.Sprintf("%s/%s/%s", payload.IpServer, payload.Path, filename)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	dir := fmt.Sprintf("cmd/video-hls-service/data/video/%s/%s", payload.Uuid, payload.Quantity)
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
