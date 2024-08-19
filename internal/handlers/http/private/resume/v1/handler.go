package handlers_http_private_resume_v1

import (
	"context"

	service_v1 "github.com/teyz/go-svc-template/internal/service/v1"
)

type Handler struct {
	service service_v1.ResumeStoreService
}

func NewHandler(_ context.Context, service service_v1.ResumeStoreService) *Handler {
	return &Handler{
		service: service,
	}
}
