package handler

import (
	"net/http"
)

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(2048)
	if err != nil {
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.UploadFile(r.MultipartForm)
	if err != nil {
		http.Error(w, "Failed to upload file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("File Uploaded Successfully"))
}
