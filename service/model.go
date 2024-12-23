package service

import videoservice "app/cmd/video-hls-service/service/video"

type Service struct {
	VideoService videoservice.VideoService
}
