package handler

import (
	"net/http"
)

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	err := h.service.DownloadFile()
	if err != nil {}
	w.Write([]byte("Download File"))
}

