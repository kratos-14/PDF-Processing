package handler

import "net/http"

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Download File"))
}
