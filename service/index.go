package service

import videoservice "app/cmd/video-hls-service/service/video"

func Register() Service {
	return Service{
		VideoService: videoservice.Register(),
	}
}
