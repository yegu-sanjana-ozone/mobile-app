package handler

import service "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/SERVICE"

type Handler struct {
	service service.Service
}

func New() *Handler{
	return &Handler{
		service: service.New(),
	}
}
