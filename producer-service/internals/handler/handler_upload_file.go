package handler

import (
	"net/http"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/model"
)

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upload file"))
	err := r.ParseMultipartForm(2048)
	if err != nil {
		err := model.Error{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(err.String())
		return
	}
	err = h.service.UploadFile(r.MultipartForm)
	if err != nil {
		err := model.Error{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(err.String())
		return
	}
}
