package server

import (
	"net/http"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/handler"
)

func New(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/files", h.UploadFile)
	mux.HandleFunc("/api/file", h.DownloadFile)
	return mux
}
