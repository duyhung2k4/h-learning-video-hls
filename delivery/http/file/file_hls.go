package filehandle

import (
	constant "app/internal/constants"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *fileHandle) FileHls(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	quantity := ctx.Param("quantity")
	filename := ctx.Param("filename")
	filepath := fmt.Sprintf("cmd/video-hls-service/data/video/%s/%s/%s", uuid, quantity, filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		logapp.Logger("get-file-hls", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	ctx.File(filepath)
}
