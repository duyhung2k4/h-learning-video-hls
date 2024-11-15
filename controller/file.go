package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type fileController struct{}

type FileController interface {
	FileHls(w http.ResponseWriter, r *http.Request)
}

func (c *fileController) FileHls(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	quantity := chi.URLParam(r, "quantity")
	filename := chi.URLParam(r, "filename")
	filepath := fmt.Sprintf("video/%s/%s/%s", uuid, quantity, filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filepath)
}

func NewFileController() FileController {
	return &fileController{}
}
