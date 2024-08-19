package handlers_http_private_resume_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	pkg_http "github.com/teyz/go-svc-template/pkg/http"
)

func (h *Handler) GetResume(c echo.Context) error {
	filePath := "./internal/handlers/http/private/resume/v1/resume.pdf"

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, c.File(filePath)))
}
