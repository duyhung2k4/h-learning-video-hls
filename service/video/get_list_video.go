package videoservice

import (
	requestdata "app/internal/dto/client"
	queuepayload "app/internal/dto/queue_payload"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (s *videoService) GetListVideo(payload queuepayload.QueueFileM3U8Payload) ([]string, error) {
	var data requestdata.GetListVideoResponse

	url := fmt.Sprintf("%s/%s", payload.IpServer, payload.Path)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal("Error fetching data: ", err)
		return []string{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
		return []string{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Error decoding JSON: ", err)
		return data.Files, err
	}

	return data.Files, nil
}
