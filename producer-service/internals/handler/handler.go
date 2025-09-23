package handler

import "github.com/kratos-14/pdf-compressor/producer-service/internals/service"

type Handler struct {
	service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}