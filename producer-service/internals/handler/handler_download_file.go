package handler

import (
	"io"
	"net/http"
)

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	fileName, writer, err := h.service.DownloadFile(id)
	if err != nil {
		http.Error(w, "Failed to download file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/pdf")
	_, err = io.Copy(w, writer)
	if err != nil {
		http.Error(w, "Failed to write file to response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
