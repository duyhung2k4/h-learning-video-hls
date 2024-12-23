package filehandle

import (
	"app/cmd/video-hls-service/service"
	constant "app/internal/constants"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type fileHandle struct {
	service service.Service
}

type FileHandle interface {
	FileHls(ctx *gin.Context)
}

func NewHandle() FileHandle {
	return &fileHandle{
		service: service.Register(),
	}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.GET_HTTP,
		Endpoint:   "video/:uuid/:quantity/:filename",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.FileHls,
	})
}
